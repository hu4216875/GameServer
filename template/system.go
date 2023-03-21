package template

type SystemItem struct {
	Id   uint32
	Para []int32
}

type SystemItemTemplate struct {
	data map[uint32]*SystemItem
}

func (i *SystemItemTemplate) load() {
	i.data = make(map[uint32]*SystemItem)
}

func (i *SystemItemTemplate) init() {
	i.data = make(map[uint32]*SystemItem)
}

func (i *SystemItemTemplate) GetSystemItem(id uint32) *SystemItem {
	return nil
}
