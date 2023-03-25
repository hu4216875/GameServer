package template

import "sync"

type Item struct {
	Id         uint32
	BigType    uint32
	SmallType  uint32
	Effect     uint32
	EffectArgs []uint32
}

type CostItem struct {
	ItemId  uint32
	ItemNum uint32
}

func NewCostItem(id, num uint32) *CostItem {
	return &CostItem{
		ItemId:  id,
		ItemNum: num,
	}
}

type ItemTemplate struct {
	IDataTempalte
	mutex sync.RWMutex
	data  map[uint32]*Item
}

func NewItemTemplate() *ItemTemplate {
	return &ItemTemplate{}
}

func (i *ItemTemplate) OnLoad() {
	i.data = make(map[uint32]*Item)
	rf := readRf(Item{})
	for k := 0; k < rf.NumRecord(); k++ {
		item := rf.Record(k).(*Item)
		i.data[item.Id] = item
	}
}

func (i *ItemTemplate) OnInit() {

}

func (i *ItemTemplate) OnReload() {
	data := make(map[uint32]*Item)
	rf := readRf(Item{})
	for k := 0; k < rf.NumRecord(); k++ {
		item := rf.Record(k).(*Item)
		data[item.Id] = item
	}

	i.mutex.Lock()
	defer i.mutex.Unlock()
	i.data = data
	i.OnInit()
}

func (i *ItemTemplate) OnDestory() {

}

// GetItem 获取道具配置
func (i *ItemTemplate) GetItem(id uint32) *Item {
	i.mutex.RLock()
	defer i.mutex.RUnlock()

	if ret, ok := i.data[id]; ok {
		return ret
	}
	return nil
}
