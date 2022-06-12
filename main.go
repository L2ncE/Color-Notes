package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"log"
	"wechat/api"
	"wechat/config"
	"wechat/dao/mongodb"
	"wechat/dao/mysql"
	"wechat/dao/redis"
	"wechat/pprof"
	"wechat/task"
)

func main() {
	task.CronInit()
	config.InitConfig()
	pprof.InitPprofMonitor()

	if err := mysql.InitGormDB(); err != nil {
		log.Printf("init gorm failed, err:%v\n", err)
	} else {
		log.Println("连接GORM成功!")
	}

	if err := redis.InitRedis(); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
	} else {
		log.Println("连接Redis成功!")
	}

	if err := mongodb.InitMongoDB(); err != nil {
		fmt.Printf("init mongo failed, err:%v\n", err)
	} else {
		log.Println("连接MongoDB成功!")
	}

	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)
	u2, err := uuid.FromString("f5394eef-e576-4709-9e4b-a7c231bd34a4")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)

	api.InitEngine()
}
