package template

import (
	"github.com/name5566/leaf/log"
	"sync"
)

type System struct {
	Id   uint32
	Para []int32
}

type SystemItemTemplate struct {
	IDataTempalte
	mutex sync.RWMutex
	data  map[uint32]*System

	oreId                   uint32      // 矿洞id
	oreSpeed                uint32      // 矿洞速度
	oreItemId               uint32      // 矿洞产出道具id
	oreUpgradeSpeedCostItem []*CostItem // 挖矿提升速度需要消耗的道具
	oreAddSpeed             uint32      // 增加的挖矿速度
}

func NewSystemItemTemplate() *SystemItemTemplate {
	return &SystemItemTemplate{}
}

func (i *SystemItemTemplate) OnLoad() {
	i.data = make(map[uint32]*System)
	rf := readRf(System{})
	for k := 0; k < rf.NumRecord(); k++ {
		sysItem := rf.Record(k).(*System)
		i.data[sysItem.Id] = sysItem
	}
}

func (i *SystemItemTemplate) OnReload() {
	data := make(map[uint32]*System)
	rf := readRf(System{})
	for k := 0; k < rf.NumRecord(); k++ {
		sysItem := rf.Record(k).(*System)
		data[sysItem.Id] = sysItem
	}

	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.data = data
	i.OnInit()
}

func (i *SystemItemTemplate) OnDestory() {

}

func (s *SystemItemTemplate) OnInit() {
	for id, data := range s.data {
		if id == uint32(SystemId_OreDistrict_Id) {
			s.oreId = uint32(data.Para[0])
		} else if id == uint32(SystemId_OreDistrict_Init_Speed) {
			s.oreSpeed = uint32(data.Para[0])
		} else if id == uint32(SystemId_OreDistrict_Make_Item_Id) {
			s.oreItemId = uint32(data.Para[0])
		} else if id == uint32(SystemId_OreDistrict_Upgrade_Speed_Cost_Item) {
			if len(data.Para)%2 != 0 {
				log.Fatal("SystemItemTemplate id:%v para not  exist", SystemId_OreDistrict_Upgrade_Speed_Cost_Item)
			}
			s.oreUpgradeSpeedCostItem = make([]*CostItem, 0, 0)
			for k := 0; k < len(data.Para); k += 2 {
				costItem := NewCostItem(uint32(data.Para[k]), uint32(data.Para[k+1]))
				s.oreUpgradeSpeedCostItem = append(s.oreUpgradeSpeedCostItem, costItem)
			}
		} else if id == uint32(SystemId_OreDistrict_Add_Speed) {
			s.oreAddSpeed = uint32(data.Para[0])
		}
	}
}

func (s *SystemItemTemplate) GetSystemItem(id uint32) *System {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if ret, ok := s.data[id]; ok {
		return ret
	}
	return nil
}

func (s *SystemItemTemplate) GetOreId() uint32 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.oreId
}

func (s *SystemItemTemplate) GetOreSpeed() uint32 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.oreSpeed
}

func (s *SystemItemTemplate) GetOreItemId() uint32 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.oreItemId
}

func (s *SystemItemTemplate) GetOreUpgradeSpeedCostItem() []*CostItem {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.oreUpgradeSpeedCostItem
}

func (s *SystemItemTemplate) GetOreAddSpeed() uint32 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.oreAddSpeed
}
