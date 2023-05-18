package tool

import (
	"strconv"

	"github.com/RorschachYang/msdm/dao"
)

func DBToJSON() {
	CardsTranslationToJson()
	LocationTranslationToJson()
	TagTranslationToJson()
}

func CardsTranslationToJson() {
	cards, _ := dao.GetAllCards()
	var cardsTranslation []CardTranslation
	for _, card := range cards {
		cid := strconv.Itoa(int(card.ID))
		newCardTranslation := CardTranslation{
			Cid:           cid,
			Name:          card.Name,
			NameZh:        card.NameZh,
			Description:   card.Description,
			DescriptionZh: card.DescriptionZh,
		}
		cardsTranslation = append(cardsTranslation, newCardTranslation)
	}
	ToJSON(cardsTranslation, "cards_zh.json", "./data/translation/")
}

func LocationTranslationToJson() {
	locations, _ := dao.GetAllLocations()
	var locationsTranslation []LocationTranslation
	for _, location := range locations {

		newLocationTranslation := LocationTranslation{
			Name:          location.Name,
			NameZh:        location.NameZh,
			Description:   location.Description,
			DescriptionZh: location.DescriptionZh,
		}
		locationsTranslation = append(locationsTranslation, newLocationTranslation)
	}
	ToJSON(locationsTranslation, "locations_zh.json", "./data/translation/")
}

func TagTranslationToJson() {
	tags, _ := dao.GetAllTags()
	var tagsTranslation []TagTranslation
	for _, tag := range tags {
		newTagTranslation := TagTranslation{
			Name:   tag.Name,
			NameZh: tag.NameZh,
		}
		tagsTranslation = append(tagsTranslation, newTagTranslation)
	}
	ToJSON(tagsTranslation, "tags_zh.json", "./data/translation/")
}
