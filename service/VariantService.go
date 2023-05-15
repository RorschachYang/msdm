package service

import (
	"strconv"

	"github.com/RorschachYang/msdm/dao"
)

func GetAllVariants() []Variant {
	variantsDO, err1 := dao.GetAllVariants()
	cardsDO, err2 := dao.GetAllCards()
	if err1 == nil && err2 == nil {
		//中文翻译不在这里更新
		var variants []Variant
		for _, variantsDO := range variantsDO {
			var artists []string
			for _, artist := range variantsDO.Artist {
				artists = append(artists, artist.Name)
			}
			var tags []Tag
			tags = []Tag{} //防止转换成json时属性值为null
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
