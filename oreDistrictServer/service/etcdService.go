package service

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/net/context"
	"log"
	"oreDistrictServer/common"
	"oreDistrictServer/conf"
	"time"
)

type ServiceRegister struct {
	// etcd client
	cli *clientv3.Client
	// service register key
	serviceKey string
	// service register prefix
	serviceKeyPrefix string
	// service register endpoint
	serverInfo string
	// leaseID
	leaseID clientv3.LeaseID
}

func NewServiceRegister(
	endpoints []string,
	serviceKeyPrefix string,
	serviceKey string,
	serverInfo string) *ServiceRegister {

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal("etcd connect faith...")
	}

	serviceReg := &ServiceRegister{
		cli:              cli,
		serviceKey:       serviceKey,
		serviceKeyPrefix: serviceKeyPrefix,
		serverInfo:       serverInfo,
	}

	return serviceReg
}

func (s *ServiceRegister) Register(ttl int64) error {
	serviceResp, err := s.cli.Get(context.Background(), s.serviceKeyPrefix, clientv3.WithPrefix())
	if err != nil {
		log.Println("etcd 操作出错")
		return err
	}
	// 首先判断当前前缀下是否有同一Lease
	for _, kv := range serviceResp.Kvs {
		// exist Grant
		if kv.Lease != 0 {
			// 设置租约ID
			s.leaseID = clientv3.LeaseID(kv.Lease)
			break
		}
	}

	// 没有租约 进行申请
	if s.leaseID == 0 {
		grant, err := s.cli.Grant(context.Background(), ttl)
		if err != nil {
			log.Println("申请租约失败...")
			return err
		}
		s.leaseID = grant.ID
	}

	// 进行注册 设置服务 并绑定租约
	_, err = s.cli.Put(context.Background(), s.serviceKey, s.serverInfo, clientv3.WithLease(s.leaseID))
	if err != nil {
		log.Println("服务注册失败...")
		return err
	}

	// 续约操作
	go s.ListenKeepAliveChan()
	return nil
}

func (s *ServiceRegister) ListenKeepAliveChan() {
	lease := clientv3.NewLease(s.cli)
	// 进行持续的续约
	keepAlive, err := lease.KeepAlive(context.Background(), s.leaseID)
	if err != nil {
		log.Fatal("keepAlive faith...")
	}

	for resp := range keepAlive {
		println("租约 :: ", resp.ID, "续约成功！")
	}
}

func registService(ttl int64) {
	endpoint := conf.Server.EtcdServerAddr
	serviceKey := common.SERVICE_PREFIX + conf.Server.RpcServerAddr
	register := NewServiceRegister(endpoint, common.SERVICE_PREFIX, serviceKey, conf.Server.RpcServerAddr)
	err := register.Register(30)
	if err != nil {
		fmt.Printf("err : %+v\n", err)
	}
}
