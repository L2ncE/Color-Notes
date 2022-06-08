package redis

import (
	"log"
	"strconv"
)

func LikeUpNoteSet(noteid int, openid string) (error, int) { //---点赞&&取消点赞
	val, err := rdb.SIsMember("like_id"+strconv.Itoa(noteid), openid).Result()
	if err != nil {
		log.Println("judge is like up error:", err)
		return err, 2
	}
	if val == false {
		_, err := rdb.SAdd("like_id"+strconv.Itoa(noteid), openid).Result()
		if err != nil {
			log.Println("set like up error:", err)
			return err, 2
		}
		return nil, 1
	} else {
		_, err := rdb.SRem("like_id"+strconv.Itoa(noteid), openid).Result()
		if err != nil {
			log.Println("set like up error:", err)
			return err, 2
		}
		return nil, 0
	}
}

func NoteLikeCount(noteid int) (int, error) { //--- 获赞次数
	val, err := rdb.SCard("like_id" + strconv.Itoa(noteid)).Result()
	if err != nil {
		log.Println("get like up count error:", err)
		return 0, err
	}
	return int(val), nil
}
