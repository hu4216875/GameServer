package common

import "time"

var (
	GS_SERVER_PREFIX             = "/server/gs/"
	GLOBAL_DB_NAME               = "global"
	GLOBAL_SERVERINFO_COLLECTION = "serverInfo"
	GLOBAL_ACCOUNT_COLLECTION    = "account"
	DB_OP_TIME_OUT               = 5 * time.Second
)
