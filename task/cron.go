package task

import (
	"github.com/robfig/cron/v3"
	"log"
	"wechat/dao/redis"
)

func CacheTask() {
	c := cron.New()
	//c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
	//c.AddFunc("CRON_TZ=Asia/Tokyo 30 11 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
	c.AddFunc("@hourly", func() {
		err := redis.MoveOpenIdToMySQL()
		if err != nil {
			log.Println("cron err", err)
		}
	})
	c.Start()
	select {}
}
