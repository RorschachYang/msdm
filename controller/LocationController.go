package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RorschachYang/msdm/service"
)

func ListLocations(w http.ResponseWriter, r *http.Request) {
	// 设置 HTTP 响应头
	w.Header().Set("Content-Type", "application/json")

	locationsCache := service.GetAllLocations()

	// 返回 JSON 数据
	if err := json.NewEncoder(w).Encode(locationsCache); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	// 解析ID参数
	defId := r.URL.Query().Get("defId")

	var foundLocation service.Location
	locationsCache := service.GetAllLocations()
	for _, location := range locationsCache {
		if location.DefID == defId {
			foundLocation = location
		}
	}

	w.Header().Set("Content-Type", "application/json")
	// 返回 JSON 数据
	if err := json.NewEncoder(w).Encode(foundLocation); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
