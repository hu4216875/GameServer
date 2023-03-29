package base

import (
	"context"
	"github.com/name5566/leaf/log"
	clientv3 "go.etcd.io/etcd/client/v3"
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
	endpoint string
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
		endpoint:         serverInfo,
	}

	return serviceReg
}

func (s *ServiceRegister) Register(ttl int64) error {
	serviceResp, err := s.cli.Get(context.Background(), s.serviceKeyPrefix, clientv3.WithPrefix())
	if err != nil {
		log.Error("etcd err:%v", err)
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
			log.Error("申请租约失败 err:%v", err)
			return err
		}
		s.leaseID = grant.ID
	}

	// 进行注册 设置服务 并绑定租约
	if _, err = s.cli.Put(context.Background(), s.serviceKey, s.endpoint, clientv3.WithLease(s.leaseID)); err != nil {
		log.Error("服务注册失败 err:%v", err)
		return err
	}

	// 续约操作
	go s.ListenKeepAliveChan()
	return nil
}

func (s *ServiceRegister) ListenKeepAliveChan() {
	lease := clientv3.NewLease(s.cli)
	keepAlive, err := lease.KeepAlive(context.Background(), s.leaseID)
	if err != nil {
		log.Fatal("keepAlive faith...")
	}
	go func() {
		for {
			<-keepAlive
		}
	}()
}
