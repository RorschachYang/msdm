package service

import (
	"regexp"
	"strconv"
	"time"

	"github.com/RorschachYang/msdm/dao"
)

func CreateDeck(name string, description string, code string, openid string) {
	// 提取code内括号后的内容并添加到切片
	re := regexp.MustCompile(`#\s*\(\d+\)\s*(.*)`)
	matches := re.FindAllStringSubmatch(code, -1)
	var cardNames []string
	for _, match := range matches {
		cardNames = append(cardNames, match[1])
	}

	cards, _ := dao.GetCardsByNames(cardNames)
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
		var cards []string
		for _, card := range deckDO.Cards {
			cardid := strconv.Itoa(int(card.ID))
			cards = append(cards, cardid)
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
