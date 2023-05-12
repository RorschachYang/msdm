package service

import (
	"strconv"

	"github.com/RorschachYang/msdm/dao"
)

func GetAllVariants() []Variant {
	variantsDO, err := dao.GetAllVariants()
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
			variant := Variant{
				Name:               variantsDO.Name,
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
