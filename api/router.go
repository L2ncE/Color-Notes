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

	err := engine.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		log.Printf("init error:%v\n", err)
		return
	}
}
