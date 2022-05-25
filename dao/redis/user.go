package redis

import (
	"fmt"
)

func IsRegister(userID int) { //当前用户是否注册
	val, err := rdb.SIsMember("openid", userID).Result()
	if err != nil {
		panic(err)
	}
	if val == false {
		fmt.Println("user don't like it")
	}

}
