package controller

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/RorschachYang/msdm/service"
)

func CreateDeck(w http.ResponseWriter, r *http.Request) {
	var deck service.Deck
	err := json.NewDecoder(r.Body).Decode(&deck)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	decodedDescription, err1 := base64.StdEncoding.DecodeString(deck.Description)
	if err1 != nil {
		http.Error(w, "Failed to decode code", http.StatusBadRequest)
		return
	}
	decodedCode, err2 := base64.StdEncoding.DecodeString(deck.Code)
	if err2 != nil {
		http.Error(w, "Failed to decode code", http.StatusBadRequest)
		return
	}
	service.CreateDeck(deck.Name, string(decodedDescription), string(decodedCode), deck.AuthorID)

	// 返回创建成功的响应
	w.WriteHeader(http.StatusCreated)
}
