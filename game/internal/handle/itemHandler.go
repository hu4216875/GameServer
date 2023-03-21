package handle

import (
	"server/game/internal/service"
	"server/publicconst"
)

func UseItemHandle() {
	service.ServMgr.GetItemService(publicconst.ItemService)
}
