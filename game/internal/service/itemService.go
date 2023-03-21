package service

type ItemService struct {
	IService
}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (i *ItemService) OnInit() {
}

func (i *ItemService) OnDestory() {

}

func (i *ItemService) AddItem(accountId int64, itemId, num uint32) {

}

func (i *ItemService) CostItem(accountId int64, itemId, num uint32) {

}

func (i *ItemService) EnoughItem(accountId int64, itemId, num uint32) {

}
