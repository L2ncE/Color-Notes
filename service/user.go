package service

import (
	"gorm.io/gorm"
	"wechat/dao/mysql"
	"wechat/dao/redis"
)

func RegisterUser(openId string) error {
	flag, err := redis.Ping()
	if flag {
		err = redis.AddOpenId(openId)
		return err
	}
	err = mysql.InsertOpenId(openId)
	return err
}

func IsRepeatOpenId(openId string) (bool, error) {
	flag, err := redis.IsOpenIdCache()
	if err != nil {
		return true, err
	}
	if flag { //有缓存
		flag, err = redis.IsRegister(openId) //查看是否注册过
		if err != nil {
			return true, err
		}
		return flag, nil
	}
	err = mysql.SelectUserByOpenId(openId)

	if err != nil {
		if err == gorm.ErrRecordNotFound { //找不到会报这个错误捏
			return false, nil
		}
		return false, err
	}
	return true, nil
}
