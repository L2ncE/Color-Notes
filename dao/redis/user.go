package redis

import (
	"log"
	"time"
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
	err = rdb.SAdd("openid_cache", id).Err()
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
	defer func() {
		go func() {
			//延时3秒执行
			time.Sleep(time.Second * 3)
			rdb.Del("openid_cache")
		}()
	}()
	set, err := rdb.SMembers("openid_cache").Result()
	if err != nil {
		log.Println("len of openid get error:", err)
		return err
	}
	for i, v := range set {
		err := mysql.InsertOpenId(v)
		if err != nil {
			log.Println("insert err:", err)
			break
		}
		log.Printf("moved %d info", i)
	}
	if err != nil {
		log.Println("move error:", err)
		return err
	}
	log.Println("move success")
	return nil
}

func Ping() (bool, error) {
	_, err := rdb.Ping().Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
