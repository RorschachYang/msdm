package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RorschachYang/msdm/service"
)

func ListVariants(w http.ResponseWriter, r *http.Request) {
	// 设置 HTTP 响应头
	w.Header().Set("Content-Type", "application/json")

	variantsCache := service.GetAllVariants()

	// 返回 JSON 数据
	if err := json.NewEncoder(w).Encode(variantsCache); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetVariant(w http.ResponseWriter, r *http.Request) {
	// 解析ID参数
	id := r.URL.Query().Get("id")

	var foundVariant service.Variant
	variantsCache := service.GetAllVariants()
	for _, variant := range variantsCache {
		if variant.Cid == id {
			foundVariant = variant
		}
	}

	w.Header().Set("Content-Type", "application/json")
	// 返回 JSON 数据
	if err := json.NewEncoder(w).Encode(foundVariant); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
