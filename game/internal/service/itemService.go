package service

import (
	"github.com/name5566/leaf/log"
	"server/game/internal/common"
	"server/game/internal/dao"
	"server/game/internal/model"
	"server/msg"
	"server/publicconst"
)

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

// LoadItems 加载道具
func (i *ItemService) LoadItems(accountId int64) (msg.ErrCode, []*model.Item) {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	if playerData == nil {
		log.Error("Account:%v AddItem player not find", accountId)
		return msg.ErrCode_SYSTEM_ERROR, nil
	}
	err, accountItem := dao.ItemDao.LoadItem(accountId)
	playerData.Items = accountItem.Items
	return err, accountItem.Items
}

// AddItem 增加道具
func (i *ItemService) AddItem(accountId int64, itemId, num uint32, source publicconst.ItemSource) {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	if playerData == nil {
		log.Error("Account:%v AddItem player not find", accountId)
		return
	}

	item := i.findItem(playerData, itemId)

	if item != nil {
		item.Num += num
		dao.ItemDao.UpdateItem(accountId, item)
		i.addItemEvent(accountId, itemId, num, source, item)
	} else {
		item = model.NewItem(itemId, num, 0)
		playerData.Items = append(playerData.Items, item)
		dao.ItemDao.AddItem(accountId, item)
		i.addItemEvent(accountId, itemId, num, source, item)
	}
}

// CostItem 扣除道具
func (i *ItemService) CostItem(accountId int64, itemId, num uint32, source publicconst.ItemSource) {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	if playerData == nil {
		log.Error("Account:%v CostItem player not find", accountId)
		return
	}

	item := i.findItem(playerData, itemId)
	if item == nil {
		log.Error("Account:%v CostItem itemId:%v not exist", accountId, itemId)
		return
	}

	if item.Num < num {
		log.Error("Account:%v CostItem itemId:%v num not enough [%v-%v]", accountId, itemId, item.Num, num)
		return
	}
	item.Num -= num
	dao.ItemDao.UpdateItem(accountId, item)
	i.costItemEvent(accountId, itemId, num, source, item)
}

// EnoughItem 是否有道具
func (i *ItemService) EnoughItem(accountId int64, itemId, num uint32) bool {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	if playerData == nil {
		log.Error("Account:%v EnoughItem player not find", accountId)
		return false
	}

	item := i.findItem(playerData, itemId)
	if item == nil {
		return false
	}

	if item.Num < num {
		return false
	}
	return true
}

func ToProtocolItem(item *model.Item) *msg.Item {
	return &msg.Item{
		ItemId:    item.Id,
		ItemNum:   item.Num,
		LimitDate: item.LimitDate,
	}
}

func ToProtocolItems(items []*model.Item) []*msg.Item {
	var ret []*msg.Item
	for i := 0; i < len(items); i++ {
		ret = append(ret, &msg.Item{
			ItemId:    items[i].Id,
			ItemNum:   items[i].Num,
			LimitDate: items[i].LimitDate,
		})
	}
	return ret
}

// updateClientItemChange 通知客户端道具变化
func (i *ItemService) updateClientItemChange(accountId int64, item *model.Item) {
	if playerData := common.PlayerMgr.FindPlayerData(accountId); playerData != nil {
		res := &msg.NotifyUpdateItem{}
		res.Items = append(res.Items, ToProtocolItem(item))
		playerData.PlayerAgent.WriteMsg(res)
	}
}

// updateClientItemsChange 通知客户端道具变化
func (i *ItemService) updateClientItemsChange(accountId int64, items []*model.Item) {
	if playerData := common.PlayerMgr.FindPlayerData(accountId); playerData != nil {
		res := &msg.NotifyUpdateItem{}
		res.Items = append(res.Items, ToProtocolItems(items)...)
		playerData.PlayerAgent.WriteMsg(res)
	}
}

// findItem 查找道具
func (i *ItemService) findItem(playerData *common.PlayerData, itemId uint32) *model.Item {
	for i := 0; i < len(playerData.Items); i++ {
		if playerData.Items[i].Id == itemId {
			return playerData.Items[i]
		}
	}
	return nil
}

// addItemEvent 添加道具事件
func (i *ItemService) addItemEvent(accountId int64, itemId, num uint32, source publicconst.ItemSource, item *model.Item) {
	i.updateClientItemChange(accountId, item)

	// 记录日志
	itemlog := model.NewItemLog(accountId, itemId, int32(num), int32(source), "")
	dao.ItemLogDao.AddItemLog(itemlog)
}

// costItemEvent 扣除道具事件
func (i *ItemService) costItemEvent(accountId int64, itemId, num uint32, source publicconst.ItemSource, item *model.Item) {
	i.updateClientItemChange(accountId, item)

	// 记录日志
	itemlog := model.NewItemLog(accountId, itemId, int32(-num), int32(source), "")
	dao.ItemLogDao.AddItemLog(itemlog)
}
