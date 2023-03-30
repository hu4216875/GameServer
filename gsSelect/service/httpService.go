package service

import (
	"github.com/gin-gonic/gin"
	"gsSelect/conf"
	"net/http"
)

func StartHttpService() {
	// 开启http服务
	router := gin.Default()
	router.GET("/login/:userId", getLoginGs)
	router.GET("/regist", getRegistGs)
	router.Run(conf.Server.HttpServer)
}

type GsAddrInfo struct {
	GsAddr string
}

func getLoginGs(c *gin.Context) {
	userId := c.Param("userId")
	gsAddr := &GsAddrInfo{
		GsAddr: getAccountByUserId(userId),
	}
	c.IndentedJSON(http.StatusOK, gsAddr)
}

func getRegistGs(c *gin.Context) {
	gsAddr := &GsAddrInfo{
		GsAddr: selectOneRegist(),
	}
	c.IndentedJSON(http.StatusOK, gsAddr)
}
