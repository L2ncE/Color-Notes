package redis

import "log"

func IsRegister(userID int) bool { //当前用户是否注册
	flag, err := rdb.SIsMember("openid", userID).Result()
	if err != nil {
		log.Println("isRegister err:", err)
		return true
	}
	return flag
}
