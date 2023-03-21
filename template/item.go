package template

import "fmt"

type Item struct {
	Id         uint32
	BigType    uint32
	SmallType  uint32
	Effect     uint32
	EffectArgs []uint32
}

type ItemTemplate struct {
	data map[uint32]*Item
}

func (i *ItemTemplate) load() {
	i.data = make(map[uint32]*Item)
	rf := readRf(Item{})
	fmt.Println(rf)
}

func (i *ItemTemplate) init() {

}

func (i *ItemTemplate) GetItem(id uint32) *Item {
	return nil
}
