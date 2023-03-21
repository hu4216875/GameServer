package publicconst

import "time"

type PlayerState int
type ServiceId int

var (
	MAX_USERID_LEN = 10

	GLOBAL_DB_NAME               = "global"
	GLOBAL_ACCOUNT_COLLECTION    = "account"
	GLOBAL_SERVERINFO_COLLECTION = "serverInfo"

	LOCAL_DB_NAME            = "local"
	LOCAL_ACCOUNT_COLLECTION = "account"
	DB_OP_TIME_OUT           = 5 * time.Second
	LOG_DB_NAME              = "log"

	CLIENT_HEART_INTERVAL = 10 // 客户端心跳间隔(s)
	MAX_CLIENT_HERART_NUM = 3  // 最大心跳包数量
)

const (
	Logining PlayerState = iota // 登录中
	Online                      // 在线
	Offline                     // 离线
)

const (
	ItemService ServiceId = iota
)
