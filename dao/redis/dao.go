package redis

import (
	"github.com/go-redis/redis"
	"log"
)

var rdb *redis.Client

func InitRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = rdb.Ping().Result()
	if err != nil {
		log.Printf("连接失败：%v\n", err)
	}
	return nil
}
