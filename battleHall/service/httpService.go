package service

import (
	"battleHall/conf"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SceneInfo struct {
	SceneAddr string `json:"sceneAddr"`
	Limit     uint32 `json:"limit"`
}

func InitHttpService() {
	// 开启http服务
	router := gin.Default()
	router.GET("/listScene", listScene)
	router.PUT("/updateScene", updateScene)
	go func() {
		router.Run(conf.Server.HttpServer)
	}()
}

func listScene(c *gin.Context) {
	data := getAllScene()
	c.IndentedJSON(http.StatusOK, data)
}

func updateScene(c *gin.Context) {
	var sceneInfo SceneInfo
	if err := c.BindJSON(&sceneInfo); err != nil {
		return
	}
	UpdateSceneLimit(sceneInfo.SceneAddr, sceneInfo.Limit)
}
