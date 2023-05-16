package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RorschachYang/msdm/service"
)

func createDeckHandler(w http.ResponseWriter, r *http.Request) {
	var deck service.Deck
	err := json.NewDecoder(r.Body).Decode(&deck)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	service.CreateDeck(deck.Name, deck.Description, deck.Code, deck.AuthorID)

	// 返回创建成功的响应
	w.WriteHeader(http.StatusCreated)
}
