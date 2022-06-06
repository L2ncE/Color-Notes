package task

import (
	"github.com/robfig/cron/v3"
	"log"
	"wechat/dao/redis"
)

func CronInit() {
	c := cron.New()
	c.Start()
	_, err := c.AddFunc("@every 1h", func() {
		err := redis.MoveOpenIdToMySQL()
		if err != nil {
			log.Println("cron err", err)
			return
		}
	})
	if err != nil {
		log.Println("cron err", err)
		return
	}
}
