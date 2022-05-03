package service

import (
	"gorm.io/gorm"
	"wechat/dao"
)

func RegisterUser(openId string) error {
	err := dao.InsertOpenId(openId)
	return err
}

func IsRepeatOpenId(openId string) (bool, error) {
	err := dao.SelectUserByOpenId(openId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
