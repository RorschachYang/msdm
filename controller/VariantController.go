package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/RorschachYang/msdm/service"
)

func ListVariants(w http.ResponseWriter, r *http.Request) {

	page := r.URL.Query().Get("page")

	pageSize := r.URL.Query().Get("page_size")

	pageNum, _ := strconv.Atoi(page)

	pageSizeNum, _ := strconv.Atoi(pageSize)

	// 设置 HTTP 响应头
	w.Header().Set("Content-Type", "application/json")

	variantsCache := service.GetAllVariants()

	var i, j int

	var result []service.Variant

	if (pageNum+1)*pageSizeNum <= len(variantsCache) {
		i = pageNum*pageSizeNum + 1
		j = (pageNum + 1) * pageSizeNum

		result = variantsCache[i:j]
	} else if (pageNum+1)*pageSizeNum > len(variantsCache) && pageNum*pageSizeNum <= len(variantsCache) {
		i = pageNum*pageSizeNum + 1
		j = len(variantsCache)

		result = variantsCache[i:j]
	} else if pageNum*pageSizeNum > len(variantsCache) {
	}

	// 返回 JSON 数据
	if err := json.NewEncoder(w).Encode(result); err != nil {
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
