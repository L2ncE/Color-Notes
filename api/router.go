package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"wechat/global"
)

func InitEngine() {
	engine := gin.Default()
	engine.Use(CORS())

	engine.POST("/user/signup", getOpenId)
	engine.POST("/test", test)

	notebookGroup := engine.Group("/notebook")
	{
		notebookGroup.Use(JWTAuth)
		notebookGroup.POST("/create", createNotebook)
		notebookGroup.PUT("/update/name/:id", changeNotebookName)
		notebookGroup.PUT("/update/color/:id", changeNotebookColor)
	}

	err := engine.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		log.Printf("init error:%v\n", err)
		return
	}
}
