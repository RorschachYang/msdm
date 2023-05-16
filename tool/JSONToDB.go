package tool

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/RorschachYang/msdm/dao"

	"gorm.io/gorm"
)

func JSONToDB() {
	CardsDataToDB()
	LocationsDataToDB()
	TitlesDataToDB()
	TranslateCardToDB() //先翻译卡牌，让变体可以获取到卡牌中文名
	TranslateLocation()
	VariantsDataToDB()
	TagTranslationToJson()
}

func CardsDataToDB() {
	// 读取JSON文件数据
	data, err := ioutil.ReadFile("./data/source/cardsdata.json")
	if err != nil {
		fmt.Println("读取JSON文件失败：", err)
		return
	}

	// 解析JSON数据到结构体
	var cardsources []CardSource
	err = json.Unmarshal(data, &cardsources)
	if err != nil {
		fmt.Println("解析JSON数据失败：", err)
		return
	}

	for _, cs := range cardsources {
		power, _ := strconv.Atoi(cs.Power)
		cost, _ := strconv.Atoi(cs.Cost)
		cid, _ := strconv.ParseUint(cs.CID, 10, 64)
		//去除链接中名称前后的字符串
		imageURLName := regexp.MustCompile(`\/([^/]+)\.webp`).FindStringSubmatch(cs.Src)
		existingCard, err := dao.GetCardByID(uint(cid))
		if err == nil {
			//中文翻译不在这里更新
			existingCard.Description = cs.Description
			existingCard.Cost = cost
			existingCard.Power = power
			existingCard.Source = cs.Source

			dao.UpdateCard(existingCard)
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			newCard := &dao.Card{
				//中文翻译不在这里更新
				ID:           uint(cid),
				Name:         cs.CardName,
				Description:  cs.Description,
				Cost:         cost,
				Power:        power,
				Source:       cs.Source,
				Defid:        cs.CardDefID,
				ImageURLName: imageURLName[1],
			}

			dao.CreateCard(newCard)
		}
	}
}

func LocationsDataToDB() {
	// 读取JSON文件数据
	data, err := ioutil.ReadFile("./data/source/locationsdata.json")
	if err != nil {
		fmt.Println("读取JSON文件失败：", err)
		return
	}

	// 解析JSON数据到结构体
	var locationSources []LocationSource
	err = json.Unmarshal(data, &locationSources)
	if err != nil {
		fmt.Println("解析JSON数据失败：", err)
		return
	}

	for _, ls := range locationSources {
		//去除链接中名称前后的字符串
		imageURLName := regexp.MustCompile(`\/([^/]+)\.webp`).FindStringSubmatch(ls.ImageSrc)

		existingLocation, err := dao.GetLocationByName(ls.Name)
		if err == nil {
			//中文翻译不在这里更新
			existingLocation.Description = ls.Ability
			existingLocation.Released = ls.Status == "released"

			dao.UpdateLocation(existingLocation)
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			newLocation := &dao.Location{
				//中文翻译不在这里更新
				Name:         ls.Name,
				Description:  ls.Ability,
				DefID:        ls.CardDefID,
				ImageURLName: imageURLName[1],
				Released:     ls.Status == "released",
			}

			dao.CreateLocation(newLocation)
		}
	}
}

func TitlesDataToDB() {
	// 读取JSON文件数据
	data, err := ioutil.ReadFile("./data/source/titlesdata.json")
	if err != nil {
		fmt.Println("读取JSON文件失败：", err)
		return
	}

	// 解析JSON数据到结构体
	var TitleSources []TitleSource
	err = json.Unmarshal(data, &TitleSources)
	if err != nil {
		fmt.Println("解析JSON数据失败：", err)
		return
	}

	for _, ts := range TitleSources {

		_, err := dao.GetTitleByName(ts.Text)
		if err == nil {
			//中文翻译不在这里更新

		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			newTitle := &dao.Title{
				//中文翻译不在这里更新
				Name: ts.Text,
			}

			dao.CreateTitle(newTitle)
		}
	}
}

func TranslateCardToDB() {
	content, err := ioutil.ReadFile("./data/translation/cards_zh.json")
	if err != nil {
		panic(err)
	}

	var cardTranslations []CardTranslation
	if err := json.Unmarshal(content, &cardTranslations); err != nil {
		panic(err)
	}

	for _, translation := range cardTranslations {
		cid, _ := strconv.ParseUint(translation.Cid, 10, 64)
		card, _ := dao.GetCardByID(uint(cid))
		if err == nil {
			card.NameZh = translation.NameZh
			card.DescriptionZh = translation.DescriptionZh

			dao.UpdateCard(card)
		}
	}
}

func TranslateLocationToDB() {
	content, err := ioutil.ReadFile("./data/translation/locations_zh.json")
	if err != nil {
		panic(err)
	}

	var locationTranslations []LocationTranslation
	if err := json.Unmarshal(content, &locationTranslations); err != nil {
		panic(err)
	}

	for _, translation := range locationTranslations {
		location, _ := dao.GetLocationByName(translation.Name)
		if err == nil {
			location.NameZh = translation.NameZh
			location.DescriptionZh = translation.DescriptionZh

			dao.UpdateLocation(location)
		}
	}
}

func VariantsDataToDB() {
	// 读取JSON文件数据
	data, err := ioutil.ReadFile("./data/source/variantsdata.json")
	if err != nil {
		fmt.Println("读取JSON文件失败：", err)
		return
	}

	// 解析JSON数据到结构体
	var variantSources []VariantSource
	err = json.Unmarshal(data, &variantSources)
	if err != nil {
		fmt.Println("解析JSON数据失败：", err)
		return
	}

	for _, vs := range variantSources {
		//将tag按空格拆分并转成大写
		tags := strings.Split(vs.Tags, " ")
		for i, tag := range tags {
			tags[i] = strings.ToUpper(tag)
		}
		cid, _ := strconv.ParseUint(vs.Cid, 10, 64)
		vid, _ := strconv.ParseUint(vs.Vid, 10, 64)
		//去除链接中名称前后的字符串
		imageURLName := regexp.MustCompile(`\/([^/]+)\.webp`).FindStringSubmatch(vs.Src)

		existingVariant, err := dao.GetVariantByID(uint(vid))
		if err == nil {
			//中文翻译不在这里更新
			existingVariant.Rarity = vs.Rarity
			existingVariant.Released = vs.Status == "released"

			dao.UpdateVariant(existingVariant)
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			newVariant := &dao.Variant{
				ID:           uint(vid),
				Name:         vs.Name,
				ImageURLName: imageURLName[1],
				Rarity:       vs.Rarity,
				CardID:       uint(cid),
				Released:     vs.Status == "released",
			}
			dao.CreateVariant(newVariant)
		}

		//将Artist创建并关联
		if vs.Sketcher != "" {
			dao.UpsertArtistAndLinkWithVariant(vs.Sketcher, uint(vid))
		}
		if vs.Inker != "" {
			dao.UpsertArtistAndLinkWithVariant(vs.Inker, uint(vid))
		}
		if vs.Colorist != "" {
			dao.UpsertArtistAndLinkWithVariant(vs.Colorist, uint(vid))
		}

		for _, tag := range tags {
			if tag != "" {
				dao.UpsertTagAndLinkWithVariant(tag, uint(vid))
			}
		}
	}
}

func TranslateTagToDB() {
	content, err := ioutil.ReadFile("./data/translation/tags_zh.json")
	if err != nil {
		panic(err)
	}

	var tagTranslations []TagTranslation
	if err := json.Unmarshal(content, &tagTranslations); err != nil {
		panic(err)
	}

	for _, translation := range tagTranslations {
		tag, _ := dao.GetTagByName(translation.Name)
		if err == nil {
			tag.NameZh = translation.NameZh
			tag.Name = translation.Name

			dao.UpdateTag(tag)
		}
	}
}
