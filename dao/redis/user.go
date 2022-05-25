package redis

import "log"

func IsRegister(id string) (bool, error) { //当前用户是否注册
	flag, err := rdb.SIsMember("openid", id).Result()
	if err != nil {
		log.Println("redis judge is register err:", err)
		return false, err
	}
	return flag, nil
}

func AddOpenId(id string) error {
	err := rdb.SAdd("openid", id).Err()
	if err != nil {
		log.Println("redis add openid err", err)
		return err
	}
	return nil
}
