package service

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/RorschachYang/msdm/dao"
)

func CreateDeck(name string, description string, code string, openid string) {

	decodedStr, _ := base64.StdEncoding.DecodeString(code)

	fmt.Println(decodedStr)
	re := regexp.MustCompile(`"CardDefId":"([^"]+)"`)
	matches := re.FindAllStringSubmatch(string(decodedStr), -1)
	fmt.Println(matches)

	var results []string

	for _, match := range matches {
		results = append(results, match[1])
	}
	fmt.Println(results)
	cards, _ := dao.GetCardsByDefids(results)
	// user, err := dao.GetUserByOpenID(openid)
	// if err != nil {
	// 	return
	// }

	newDeck := &dao.Deck{
		Name:        name,
		Description: description,
		Code:        code,
		Cards:       cards,
		// Author:      *user,
	}

	dao.CreateDeck(newDeck)
}

func GetRecentlyCreatedDecks(days int) ([]Deck, error) {
	duration := time.Duration(days) * 24 * time.Hour

	decksDO, err := dao.GetRecentlyCreatedDecks(time.Duration(duration))
	if err != nil {
		return nil, err
	}

	var decks []Deck
	for _, deckDO := range decksDO {
		id := strconv.Itoa(int(deckDO.ID))
		var cards []Card
		for _, card := range deckDO.Cards {
			newCard := Card{
				Cid:                strconv.Itoa(int(card.ID)),
				Cost:               card.Cost,
				ImageURLCompressed: CompressedImageURLPrefix + "cards/" + card.ImageURLName + ".webp",
			}
			cards = append(cards, newCard)
		}
		newDeck := Deck{
			ID:          id,
			Name:        deckDO.Name,
			Description: deckDO.Description,
			CopiedTimes: deckDO.CopiedTimes,
			Code:        deckDO.Code,
			AuthorID:    deckDO.Author.OpenID,
			Cards:       cards,
		}
		decks = append(decks, newDeck)
	}

	return decks, nil
}

func DeleteDeck(id uint) error {
	err := dao.DeleteDeckByID(id)
	if err != nil {
		return err
	}
	return nil
}
