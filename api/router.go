package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitEngine() {
	engine := gin.Default()
	engine.Use(CORS())

	engine.POST("/user/signup", test)

	//err := engine.Run(fmt.Sprintf(":%d", global.Settings.Port))
	err := engine.Run(":5556")
	if err != nil {
		log.Printf("init error:%v\n", err)
		return
	}
}
