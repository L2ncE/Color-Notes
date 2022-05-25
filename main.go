package main

import (
	"fmt"
	"wechat/api"
	"wechat/config"
	"wechat/dao"
)

func main() {
	config.InitConfig()

	if err := dao.InitGormDB(); err != nil {
		fmt.Printf("init gorm failed, err:%v\n", err)
		return
	}
	fmt.Println("连接GORM MySQL数据库成功!")

	api.InitEngine()
}
