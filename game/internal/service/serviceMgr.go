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
	m.registService(publicconst.GMService, NewGmService())
	m.registService(publicconst.OreService, NewOreService())
	m.registService(publicconst.AccountService, NewAccountService())

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

func (m *ServiceMgr) GetItemService() *ItemService {
	if data, ok := m.serviceMap[publicconst.ItemService]; ok {
		return data.(*ItemService)
	}
	return nil
}

func (m *ServiceMgr) GetGmService() *GmService {
	if data, ok := m.serviceMap[publicconst.GMService]; ok {
		return data.(*GmService)
	}
	return nil
}

func (m *ServiceMgr) GetOreService() *OreService {
	if data, ok := m.serviceMap[publicconst.OreService]; ok {
		return data.(*OreService)
	}
	return nil
}

func (m *ServiceMgr) GetAccountService() *AccountService {
	if data, ok := m.serviceMap[publicconst.AccountService]; ok {
		return data.(*AccountService)
	}
	return nil
}
