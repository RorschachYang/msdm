package service

import (
	"encoding/base64"
	"encoding/json"
	"regexp"
	"strconv"
	"time"

	"github.com/RorschachYang/msdm/dao"
)

func CreateDeck(name string, description string, code string, openid string) {
	regex := regexp.MustCompile(`#([^#]+)#`)
	matches := regex.FindStringSubmatch(code)

	decodedStr, _ := base64.StdEncoding.DecodeString(matches[0])

	var data map[string][]map[string]string
	if err := json.Unmarshal([]byte(decodedStr), &data); err != nil {
		panic(err)
	}

	var cardDefIds []string
	for _, card := range data["Cards"] {
		cardDefIds = append(cardDefIds, card["CardDefId"])
	}

	cards, _ := dao.GetCardsByDefids(cardDefIds)
	user, err := dao.GetUserByOpenID(openid)
	if err != nil {
		return
	}

	newDeck := &dao.Deck{
		Name:        name,
		Description: description,
		Code:        code,
		Cards:       cards,
		Author:      *user,
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
