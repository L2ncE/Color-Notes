package service

import (
	"gorm.io/gorm"
	"wechat/dao/mysql"
)

func RegisterUser(openId string) error {
	err := mysql.InsertOpenId(openId)
	return err
}

func IsRepeatOpenId(openId string) (bool, error) {
	err := mysql.SelectUserByOpenId(openId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
