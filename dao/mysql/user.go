package mysql

import (
	"log"
	"wechat/model"
)

func InsertOpenId(openId string) error {
	deres := db.Select("OpenId").Create(&model.User{OpenId: openId})
	err := deres.Error
	if err != nil {
		log.Printf("insert failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectUserByOpenId(openId string) error {
	var user model.User
	dbRes := db.Model(&model.User{}).Select(user).Where("OpenId = ?", openId).First(&user)
	err := dbRes.Error
	if err != nil {
		log.Println("select failed, err:", err)
		return err
	}
	return nil
}
