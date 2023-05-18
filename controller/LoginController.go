package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RorschachYang/msdm/service"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// 获取小程序端发送的code
	code := r.URL.Query().Get("code")

	var openid = service.Login(code)

	// 将open_id返回给小程序
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"openid": openid,
	})
}
