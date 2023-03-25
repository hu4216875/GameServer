package template

type SystemId uint32

const (
	SystemId_None                                SystemId = iota
	SystemId_OreDistrict_Init_Speed                       // 矿区初始挖矿速度
	SystemId_OreDistrict_Make_Item_Id                     // 矿区产出的道具id
	SystemId_OreDistrict_Upgrade_Speed_Cost_Item          // 挖矿提升速度需要消耗的道具
	SystemId_OreDistrict_Id
	SystemId_OreDistrict_Add_Speed //  矿洞id
)

const (
	ITEM_FILE   = "item"
	SYSTEM_FILE = "system"
)
