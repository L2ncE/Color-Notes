package config

import (
	"github.com/spf13/viper"
	"log"
	"wechat/global"
	"wechat/model"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile("./setting-dev.yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	serverConfig := model.ServerConfig{}
	//给serverConfig初始值
	err = v.Unmarshal(&serverConfig)
	if err != nil {
		log.Println(err)
	}
	// 传递给全局变量
	global.Settings = serverConfig
}
