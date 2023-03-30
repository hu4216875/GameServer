package main

import (
	"gsSelect/common"
	"gsSelect/db"
	"gsSelect/service"
)

func main() {
	common.InitLog()
	defer common.ReleaseLog()
	db.OnInit()
	defer db.OnRelease()

	service.InitService()

	service.StartHttpService()
}
