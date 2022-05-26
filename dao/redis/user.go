package redis

import (
	"log"
	"wechat/dao/mysql"
)

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

func IsOpenIdCache() (bool, error) {
	es, err := rdb.SMembers("openid").Result()
	if err != nil {
		log.Println("openid cache get error:", err)
		return false, err
	}
	if len(es) > 0 {
		return true, nil
	}
	return false, nil
}

func MoveOpenIdToMySQL() error {
	set, err := rdb.SMembers("openid").Result()
	if err != nil {
		log.Println("len of openid get error:", err)
		return err
	}
	for _, v := range set {
		err := mysql.InsertOpenId(v)
		if err != nil {
			log.Println("insert err:", err)
			break
		}
	}
	return nil
}
