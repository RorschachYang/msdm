package service

import (
	"strconv"

	"github.com/RorschachYang/msdm/dao"
)

func GetAllCards() []Card {
	cardsDO, err := dao.GetAllCards()
	if err == nil {
		//中文翻译不在这里更新
		var cards []Card
		for _, cardsDO := range cardsDO {
			cid := strconv.FormatUint(uint64(cardsDO.ID), 10)
			card := Card{
				Name:               cardsDO.Name,
				NameZh:             cardsDO.NameZh,
				Description:        cardsDO.Description,
				DescriptionZh:      cardsDO.DescriptionZh,
				Source:             cardsDO.Source,
				Power:              cardsDO.Power,
				Cost:               cardsDO.Cost,
				ImageURL:           ImageURLPrefix + cardsDO.ImageURLName + ".webp",
				Cid:                cid,
				Defid:              cardsDO.Defid,
				ImageURLCompressed: CompressedImageURLPrefix + "cards/" + cardsDO.ImageURLName + ".webp",
			}
			cards = append(cards, card)
		}
		return cards
	}
	return nil
}
