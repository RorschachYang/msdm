package service

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RorschachYang/msdm/dao"
	"gorm.io/gorm"
)

type wxLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func Login(code string) string {

	// 构造请求URL，调用微信登录API接口
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appID, appSecret, code)

	// 创建跳过 TLS 证书验证的 Transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 创建 HTTP 客户端，并使用自定义 Transport
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 解析微信登录API响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var loginResp wxLoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		panic(err)
	}

	CreateUser(loginResp.OpenID)

	return loginResp.OpenID
}

func CreateUser(openid string) {
	// 将openid存入数据库
	existingUser, err := dao.GetUserByOpenID(openid)
	if err == nil {
		fmt.Println("User login, ID is" + string(existingUser.ID))

	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		newUser := &dao.User{
			OpenID: openid,
		}

		dao.CreateUser(newUser)
	}
}
