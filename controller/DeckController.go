package controller

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

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
	// decodedCode, err2 := base64.StdEncoding.DecodeString(deck.Code)
	// if err2 != nil {
	// 	http.Error(w, "Failed to decode code", http.StatusBadRequest)
	// 	return
	// }
	service.CreateDeck(deck.Name, string(decodedDescription), deck.Code, deck.AuthorID)

	// 返回创建成功的响应
	w.WriteHeader(http.StatusCreated)
}

func GetDecksCreatedLastDays(w http.ResponseWriter, r *http.Request) {
	// 解析参数
	d := r.URL.Query().Get("days")

	days, _ := strconv.Atoi(d)

	decks, _ := service.GetRecentlyCreatedDecks(days)

	w.Header().Set("Content-Type", "application/json")
	// 返回 JSON 数据
	if err := json.NewEncoder(w).Encode(decks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteDeck(w http.ResponseWriter, r *http.Request) {
	// 解析ID参数
	idstr := r.URL.Query().Get("id")

	id, _ := strconv.Atoi(idstr)

	service.DeleteDeck(uint(id))

	// 返回创建成功的响应
	w.WriteHeader(http.StatusCreated)
}

func GetDeckByID(w http.ResponseWriter, r *http.Request) {
	// 解析ID参数
	id := r.URL.Query().Get("id")

	idstr, _ := strconv.Atoi(id)

	deck := service.GetDeckByID(uint(idstr))

	// 返回 JSON 数据
	if err := json.NewEncoder(w).Encode(&deck); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUpers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uperlist := service.GetUpers()

	// jsonData, _ := json.Marshal(uperlist)
	// 返回 JSON 数据
	if err := json.NewEncoder(w).Encode(uperlist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
