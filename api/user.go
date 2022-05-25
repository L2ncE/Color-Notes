package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq/v2"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"wechat/model"
	"wechat/service"
	"wechat/util"
)

//https://api.weixin.qq.com/sns/jscode2session?appid=wx909f50d56919e970&secret=dd7cdc3c91c868e3b69fbc497a664d4d&js_code=023hDL100VmJMN1KCQ100JkoWf0hDL1A&grant_type=authorization_code
const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wx811301ff5d288f81"
	appSecret       = "a66357dc0589ac756030bcdbc641c410"
)

func getOpenId(c *gin.Context) {

	//传code
	code := c.PostForm("code")

	//调用auth.code2Session接口获取openid
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, code)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("get openid error", err)
		util.RespError(c, 400, "get openid error")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read resp error", err)
		util.RespError(c, 400, "read resp error")
		return
	}
	json := gojsonq.New().FromString(string(body)).Find("openid")
	openId := json.(string)

	flag, err := service.IsRepeatOpenId(openId)

	if flag {
		//如果是已注册的OpenId就直接返回值
		JWT(c, openId)
		return
	}

	err = service.RegisterUser(openId)
	if err != nil {
		log.Println("register err:", err)
		util.RespError(c, 400, "register err")
		return
	}

	JWT(c, openId)
}

func JWT(c *gin.Context, openid string) {
	claim := model.MyClaims{
		OpenId: openid,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 2592000, //30天，仅做测试
			Issuer:    "YuanXinHao",
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		util.RespError(c, 400, err)
	}
	util.RespSuccessfulWithData(c, "get token successful", s)
}
