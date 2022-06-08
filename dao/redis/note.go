package redis

import (
	"log"
	"strconv"
)

func StoreUpNoteSet(noteid int, openid string) (error, int) { //---点赞&&取消点赞
	val, err := rdb.SIsMember("store_id"+strconv.Itoa(noteid), openid).Result()
	if err != nil {
		log.Println("judge is store up error:", err)
		return err, 2
	}
	if val == false {
		_, err := rdb.SAdd("store_id"+strconv.Itoa(noteid), openid).Result()
		if err != nil {
			log.Println("set store up error:", err)
			return err, 2
		}
		return nil, 1
	} else {
		_, err := rdb.SRem("store_id"+strconv.Itoa(noteid), openid).Result()
		if err != nil {
			log.Println("set store up error:", err)
			return err, 2
		}
		return nil, 0
	}
}

func NoteStoreCount(noteid int) (int, error) { //--- 获赞次数
	val, err := rdb.SCard("store_id" + strconv.Itoa(noteid)).Result()
	if err != nil {
		log.Println("get store up count error:", err)
		return 0, err
	}
	return int(val), nil
}
