package service

import (
	"strconv"

	"github.com/RorschachYang/msdm/dao"
)

func GetAllVariants() []Variant {
	variantsDO, err := dao.GetAllVariants()
	cardsDO, err := dao.GetAllCards()
	if err == nil {
		//中文翻译不在这里更新
		var variants []Variant
		for _, variantsDO := range variantsDO {
			var artists []string
			for _, artist := range variantsDO.Artist {
				artists = append(artists, artist.Name)
			}
			var tags []Tag
			for _, tagDO := range variantsDO.Tags {
				tag := Tag{Name: tagDO.Name, NameZh: tagDO.NameZh}
				tags = append(tags, tag)
			}
			vid := strconv.FormatUint(uint64(variantsDO.ID), 10)
			cid := strconv.FormatUint(uint64(variantsDO.CardID), 10)
			var cardNameZh string
			for _, cardDO := range cardsDO {
				if cardDO.ID == variantsDO.CardID {
					cardNameZh = cardDO.NameZh
				}

			}
			variant := Variant{
				Name:               variantsDO.Name,
				NameZh:             cardNameZh,
				Rarity:             variantsDO.Rarity,
				Vid:                vid,
				Released:           variantsDO.Released,
				ImageURL:           ImageURLPrefix + variantsDO.ImageURLName + ".webp",
				ImageURLCompressed: CompressedImageURLPrefix + "variants/" + variantsDO.ImageURLName + ".webp",
				Artist:             artists,
				Tags:               tags,
				Cid:                cid,
			}
			variants = append(variants, variant)
		}
		return variants
	}
	return nil
}
