package service

import (
	"server/publicconst"
)

var (
	ServMgr ServiceMgr
)

type IService interface {
	OnInit()
	OnDestory()
}

type ServiceMgr struct {
	serviceMap map[publicconst.ServiceId]IService
}

func (m *ServiceMgr) InitService() {
	m.serviceMap = make(map[publicconst.ServiceId]IService)
	m.registService(publicconst.ItemService, NewItemService())

	for _, service := range m.serviceMap {
		service.OnInit()
	}
}

func (m *ServiceMgr) Destory() {
	for _, service := range m.serviceMap {
		service.OnDestory()
	}
}

func (m *ServiceMgr) registService(serviceId publicconst.ServiceId, service IService) {
	m.serviceMap[serviceId] = service
}

func (m *ServiceMgr) GetItemService(serviceId publicconst.ServiceId) *ItemService {
	if data, ok := m.serviceMap[serviceId]; ok {
		return data.(*ItemService)
	}
	return nil
}
