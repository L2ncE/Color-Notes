package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq/v2"
	"io/ioutil"
	"net/http"
	"wechat/service"
	tool "wechat/util"
)

const (
	code2sessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"
	appID           = "wx909f50d56919e970"
	appSecret       = "f889c8e4d1609017d71b39ab90630366"
)

func getOpenId(c *gin.Context) {

	//获取code
	code := c.PostForm("code")

	//调用auth.code2Session接口获取openid
	url := fmt.Sprintf(code2sessionURL, appID, appSecret, code)
	resp, err := http.Get(url)
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		tool.RespInternalError(c)
		return
	}
	json := gojsonq.New().FromString(string(body)).Find("openid")
	openId := json.(string)

	flag, err := service.IsRepeatOpenId(openId)
	if err != nil {
		tool.RespInternalError(c)
		return
	}

	if flag {
		//如果是已注册的OpenId就直接返回值
		tool.RespSuccessfulWithData(c, openId)
		return
	}

	err = service.RegisterUser(openId)
	if err != nil {
		tool.RespInternalError(c)
		return
	}

	tool.RespSuccessfulWithData(c, openId)
	return
}