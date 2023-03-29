package common

import "time"

var (
	MAX_SERVICE_TIME        = 10 // 续约时间
	SERVICE_PREFIX          = "/server/rpcServer/hallServer/"
	DB_BATTLE_HALL          = "battleHall"
	DB_BATTLE_HALL_SERVER   = "battleServer"
	DB_OP_TIME_OUT          = 10 * time.Second
	RPC_SCENE_SERVER_PREFIX = "/server/rpcServer/sceneServer/"
	RPC_GS_SERVER_PREFIX    = "/server/gsServer/"

	MAX_SCENE_LIMIT = 1000
)
