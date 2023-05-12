package controller

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var appID string
var appSecret string

func init() {
	appID = os.Getenv("APP_ID")
	appSecret = os.Getenv("APP_SECRET")
}

type wxLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	// 获取小程序端发送的code
	code := r.URL.Query().Get("code")

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

	// 将open_id返回给小程序
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"openid": loginResp.OpenID,
	})
}
