package publicconst

import "time"

type PlayerState int
type ServiceId int
type ItemSource int

var (
	MAX_USERID_LEN = 10

	GLOBAL_DB_NAME               = "global"
	GLOBAL_ACCOUNT_COLLECTION    = "account"
	GLOBAL_SERVERINFO_COLLECTION = "serverInfo"
	LOG_DB_NAME                  = "log"

	LOCAL_DB_NAME            = "local"
	LOCAL_ACCOUNT_COLLECTION = "account"
	LOACL_Item               = "item"
	DB_OP_TIME_OUT           = 20 * time.Second

	CLIENT_HEART_INTERVAL = 10 // 客户端心跳间隔(s)
	MAX_CLIENT_HERART_NUM = 3  // 最大心跳包数量

	MAX_UPDATE_ORE_TOTAL_TIME = 10 // 更新矿洞总量时间

	MONGO_NO_RESULT = "mongo: no documents in result"

	MAX_RECYCLE_PLAYER_DATA = 3600 // 玩家数据保留1小时

	REFRESH_ORE_INTEVAL = 10 // 刷新矿洞总量间隔
)

const (
	Logining PlayerState = iota // 登录中
	Online                      // 在线
	Offline                     // 离线
)

const (
	GMService ServiceId = iota
	ItemService
	OreService
	AccountService
)

const (
	OreAddItem ItemSource = iota // 挖矿获得
	GMAddItem                    // gm 获得
	OreUpgradeSpeed
)
