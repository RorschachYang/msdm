package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RorschachYang/msdm/service"
)

func ListCards(w http.ResponseWriter, r *http.Request) {
	// 设置 HTTP 响应头
	w.Header().Set("Content-Type", "application/json")

	cardsCache := service.GetCardsCache()

	// 返回 JSON 数据
	if err := json.NewEncoder(w).Encode(cardsCache); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetCard(w http.ResponseWriter, r *http.Request) {
	// 解析ID参数
	id := r.URL.Query().Get("id")

	var foundCard service.Card
	cardsCache := service.GetCardsCache()
	for _, card := range cardsCache {
		if card.Cid == id {
			foundCard = card
		}
	}

	w.Header().Set("Content-Type", "application/json")
	// 返回 JSON 数据
	if err := json.NewEncoder(w).Encode(foundCard); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
