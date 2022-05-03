package main

import (
	"fmt"
	"wechat/api"
	"wechat/dao"
)

func main() {
	err := dao.InitGormDB()
	if err != nil {
		fmt.Printf("init failed, err:%v\n", err)
	} else {
		fmt.Println("连接GORM MySQL数据库成功!")
	}
	api.InitEngine()
}
