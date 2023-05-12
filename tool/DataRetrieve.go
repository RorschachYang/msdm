package tool

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ConvertCardsData() {
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

	// 输出解析结果
	fmt.Println("解析结果：", cardsources)

	var cards []Card
	for _, cs := range cardsources {
		power, _ := strconv.Atoi(cs.Power)
		cost, _ := strconv.Atoi(cs.Cost)
		//去除链接后的参数
		imageURL := regexp.MustCompile(`\.webp\?v=\d+`).ReplaceAllString(cs.Src, ".webp")

		card := Card{
			Name:        cs.CardName,
			Description: cs.Description,
			Source:      cs.Source,
			Power:       power,
			Cost:        cost,
			ImageURL:    imageURL,
			Cid:         cs.CID,
			Defid:       cs.CardDefID,
			Variants:    findVariants(cs.CID),
			ImageURLCompressed: strings.Replace(imageURL,
				"https://marvelsnapzone.com/wp-content/themes/blocksy-child/assets/media/cards/",
				"https://prod-7gwp0o796e119105-1317452426.tcloudbaseapp.com/assets_compressed/cards/",
				1),
		}

		cards = append(cards, card)
	}

	ToJSON(cards, "cards.json", "./data/data/")
}

func ConvertVariantsData() {
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

	// 输出解析结果
	fmt.Println("解析结果：", variantSources)

	var variants []Variant
	for _, vs := range variantSources {
		//将tag按空格拆分并转成大写
		tags := strings.Split(vs.Tags, " ")
		for i, tag := range tags {
			tags[i] = strings.ToUpper(tag)
		}
		//去除链接后的参数
		imageURL := regexp.MustCompile(`\.webp\?v=\d+`).ReplaceAllString(vs.Src, ".webp")

		variant := Variant{
			Name:     vs.Name,
			ImageURL: imageURL,
			Artist:   vs.Sketcher,
			Inker:    vs.Inker,
			Colorist: vs.Colorist,
			Tags:     tags,
			Rarity:   vs.Rarity,
			Cid:      vs.Cid,
			Vid:      vs.Vid,
			Released: vs.Status == "released",
			ImageURLCompressed: strings.Replace(imageURL,
				"https://marvelsnapzone.com/wp-content/themes/blocksy-child/assets/media/cards/",
				"https://prod-7gwp0o796e119105-1317452426.tcloudbaseapp.com/assets_compressed/variants/",
				1),
		}

		variants = append(variants, variant)
	}

	ToJSON(variants, "variants.json", "./data/data/")
}

func ConvertLocationsData() {
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

	// 输出解析结果
	fmt.Println("解析结果：", locationSources)

	var locations []Location
	for _, ls := range locationSources {
		//去除链接后的参数
		imageURL := regexp.MustCompile(`\.webp\?v=\d+`).ReplaceAllString(ls.ImageSrc, ".webp")

		location := Location{
			Name:        ls.Name,
			Description: ls.Ability,
			ImageURL:    imageURL,
			DefID:       ls.CardDefID,
			Released:    ls.Status == "released",
			ImageURLCompressed: strings.Replace(imageURL,
				"https://marvelsnapzone.com/wp-content/themes/blocksy-child/assets/media/cards/",
				"https://prod-7gwp0o796e119105-1317452426.tcloudbaseapp.com/assets_compressed/locations/",
				1),
		}

		locations = append(locations, location)
	}

	ToJSON(locations, "locations.json", "./data/data/")
}

func findVariants(cid string) []Variant {
	// 读取JSON文件数据
	data, err := ioutil.ReadFile("./data/data/variants.json")
	if err != nil {
		fmt.Println("读取JSON文件失败：", err)
	}

	// 解析JSON数据到结构体
	var variants []Variant
	err = json.Unmarshal(data, &variants)
	if err != nil {
		fmt.Println("解析JSON数据失败：", err)
	}

	var variantsFound []Variant
	for _, v := range variants {
		if v.Cid == cid {
			variantsFound = append(variantsFound, v)
		}
	}

	return variantsFound
}

func ConvertVariantsTags() {
	// 读取JSON文件数据
	data, err := ioutil.ReadFile("./data/data/variants.json")
	if err != nil {
		fmt.Println("读取JSON文件失败：", err)
	}

	// 解析JSON数据到结构体
	var variants []Variant
	err = json.Unmarshal(data, &variants)
	if err != nil {
		fmt.Println("解析JSON数据失败：", err)
	}

	// 记录所有tag并去重
	tagMap := make(map[string]bool)
	for _, variant := range variants {
		for _, tag := range variant.Tags {
			tagMap[tag] = true
		}
	}

	tags := make([]string, 0, len(tagMap))
	for tag := range tagMap {
		tags = append(tags, tag)
	}

	var variantTags []VariantTag
	for _, tag := range tags {
		vt := VariantTag{
			Value: tag,
		}

		variantTags = append(variantTags, vt)
	}

	ToJSON(variantTags, "variantTags.json", "./data/data/")
}

type CardTranslation struct {
	Cid           string `json:"cid"`
	Description   string `json:"description"`
	DescriptionZh string `json:"descriptionZh"`
	NameZh        string `json:"nameZh"`
	Name          string `json:"name"`
}

type LocationTranslation struct {
	Name          string `json:"name"`
	NameZh        string `json:"nameZh"`
	Description   string `json:"description"`
	DescriptionZh string `json:"descriptionZh"`
	DefID         string `json:"defId"`
}

func TranslateCard() {
	// 读取第一个文件
	content, err := ioutil.ReadFile("./data/data/cards.json")
	if err != nil {
		panic(err)
	}

	// 解析第一个文件中的多个卡牌
	var cards []Card
	if err := json.Unmarshal(content, &cards); err != nil {
		panic(err)
	}

	// 读取第二个文件
	content, err = ioutil.ReadFile("./data/translation/cards_zh.json")
	if err != nil {
		panic(err)
	}

	// 解析第二个文件中的多个翻译
	var cardTranslations []CardTranslation
	if err := json.Unmarshal(content, &cardTranslations); err != nil {
		panic(err)
	}

	// 将翻译内容添加到第一个文件中对应的卡牌中
	for i := range cards {
		for _, translation := range cardTranslations {
			if cards[i].Cid == translation.Cid {
				cards[i].NameZh = translation.NameZh
				cards[i].DescriptionZh = translation.DescriptionZh
				break
			}
		}
	}

	// 将结果编码成 JSON 格式输出
	ToJSON(cards, "cards.json", "./data/data/")

}

func TranslateLocation() {
	// 读取第一个文件
	content, err := ioutil.ReadFile("./data/data/locations.json")
	if err != nil {
		panic(err)
	}

	// 解析第一个文件中的多个卡牌
	var locations []Location
	if err := json.Unmarshal(content, &locations); err != nil {
		panic(err)
	}

	// 读取第二个文件
	content, err = ioutil.ReadFile("./data/translation/locations_zh.json")
	if err != nil {
		panic(err)
	}

	// 解析第二个文件中的多个翻译
	var locationTranslations []LocationTranslation
	if err := json.Unmarshal(content, &locationTranslations); err != nil {
		panic(err)
	}

	// 将翻译内容添加到第一个文件中对应的卡牌中
	for i := range locations {
		for _, translation := range locationTranslations {
			if locations[i].DefID == translation.DefID {
				locations[i].NameZh = translation.NameZh
				locations[i].DescriptionZh = translation.DescriptionZh
				break
			}
		}
	}

	// 将结果编码成 JSON 格式输出
	ToJSON(locations, "locations.json", "./data/data/")

}

func TranslateVariant() {
	// 读取第一个文件
	content, err := ioutil.ReadFile("./data/data/variants.json")
	if err != nil {
		panic(err)
	}

	// 解析第一个文件中的多个卡牌
	var variants []Variant
	if err := json.Unmarshal(content, &variants); err != nil {
		panic(err)
	}

	// 读取第二个文件
	content, err = ioutil.ReadFile("./data/translation/cards_zh.json")
	if err != nil {
		panic(err)
	}

	// 解析第二个文件中的多个翻译
	var cardTranslations []CardTranslation
	if err := json.Unmarshal(content, &cardTranslations); err != nil {
		panic(err)
	}

	// 将翻译内容添加到第一个文件中对应的卡牌中
	for i := range variants {
		for _, translation := range cardTranslations {
			if variants[i].Cid == translation.Cid {
				variants[i].NameZh = translation.NameZh
				break
			}
		}
	}

	// 将结果编码成 JSON 格式输出
	ToJSON(variants, "variants.json", "./data/data/")

}

func GetCardsFromHTML() {
	// Read HTML file
	html, err := ioutil.ReadFile("./data/html/cards-list.html")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTML file read successfully")

	// Load HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(html)))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTML document loaded successfully")

	// Find all cards and create Card objects
	var cards []CardSource
	doc.Find("a.simple-card").Each(func(i int, s *goquery.Selection) {
		var card CardSource
		card.Name = s.AttrOr("data-name", "")
		card.Cost = s.AttrOr("data-cost", "")
		card.Power = s.AttrOr("data-power", "")
		card.Ability = s.AttrOr("data-ability", "")
		card.Src = s.AttrOr("data-src", "")
		card.CID = s.AttrOr("data-cid", "")
		card.Rarity = s.AttrOr("data-rarity", "")
		card.Sketcher = s.AttrOr("data-sketcher", "")
		card.Colorist = s.AttrOr("data-colorist", "")
		card.Inker = s.AttrOr("data-inker", "")
		card.Status = s.AttrOr("data-status", "")
		card.CardDefID = s.AttrOr("data-carddefid", "")
		card.Source = s.AttrOr("data-source", "")
		card.Tags = s.AttrOr("data-tags", "")
		card.HRef = s.AttrOr("href", "")
		card.CardName = s.Find("div.cardname").Text()

		// Find all variants
		s.Find("img.lazy").Each(func(j int, v *goquery.Selection) {
			card.Variants = append(card.Variants, v.AttrOr("data-src", ""))
		})

		str := s.Find(".card-description").Text()
		re := regexp.MustCompile(`\n\s*`)               // 匹配换行符及其后的连续空格
		card.Description = re.ReplaceAllString(str, "") // 将匹配到的内容替换为空字符串

		cards = append(cards, card)
	})
	log.Println("All cards processed successfully")

	// Convert to JSON and write to file
	ToJSON(cards, "cardsdata.json", "./data/source/")
}

func GetVariantsFromHTML() {
	// Read HTML file
	html, err := ioutil.ReadFile("./data/html/variants-list.html")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTML file read successfully")

	// Load HTML document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(html)))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("HTML document loaded successfully")

	var cards []VariantSource

	doc.Find("a.simple-card").Each(func(i int, s *goquery.Selection) {
		card := VariantSource{
			Name:        strings.TrimSpace(s.Find(".cardname").Text()),
			Cost:        s.AttrOr("data-cost", ""),
			Power:       s.AttrOr("data-power", ""),
			Ability:     strings.TrimSpace(s.AttrOr("data-ability", "")),
			Src:         s.AttrOr("data-src", ""),
			Cid:         s.AttrOr("data-cid", ""),
			Vid:         s.AttrOr("data-vid", ""),
			Rarity:      s.AttrOr("data-rarity", ""),
			Sketcher:    s.AttrOr("data-sketcher", ""),
			Colorist:    s.AttrOr("data-colorist", ""),
			Inker:       s.AttrOr("data-inker", ""),
			Status:      s.AttrOr("data-status", ""),
			CardDefID:   s.AttrOr("data-carddefid", ""),
			Source:      s.AttrOr("data-source", ""),
			Tags:        s.AttrOr("data-tags", ""),
			Description: strings.TrimSpace(s.Find(".card-description").Text()),
			HRef:        s.AttrOr("href", ""),
		}
		cards = append(cards, card)
	})

	ToJSON(cards, "variantsdata.json", "./data/source/")
}

func GetArtistsFromHTML() {
	files := []string{
		"./data/html/artists-list.html",
		"./data/html/colorists-list.html",
		"./data/html/inkers-list.html",
	}

	people := make(map[string]ArtistSource)

	for _, file := range files {
		html, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(html)))
		if err != nil {
			panic(err)
		}

		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			name := strings.TrimSpace(s.Text())

			if _, ok := people[name]; !ok {
				people[name] = ArtistSource{Name: name, URL: url}
			}
		})
	}

	peopleSlice := make([]ArtistSource, 0, len(people))
	for _, person := range people {
		peopleSlice = append(peopleSlice, person)
	}

	ToJSON(peopleSlice, "artistsdata.json", "./data/source/")
}

func GetLocationsFromHTML() {
	// 读取HTML文件
	file, err := ioutil.ReadFile("./data/html/locations-list.html")
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(file)))
	if err != nil {
		panic(err)
	}

	// 遍历每个location标签，并解析属性
	locations := make([]LocationSource, 0)
	doc.Find("a.location").Each(func(i int, s *goquery.Selection) {
		location := LocationSource{
			Name:       s.AttrOr("data-name", ""),
			Cost:       s.AttrOr("data-cost", ""),
			Power:      s.AttrOr("data-power", ""),
			Ability:    s.AttrOr("data-ability", ""),
			ImageSrc:   s.AttrOr("data-src", ""),
			Rarity:     s.AttrOr("data-rarity", ""),
			Difficulty: s.AttrOr("data-difficulty", ""),
			Status:     s.AttrOr("data-status", ""),
			CardDefID:  s.AttrOr("data-carddefid", ""),
			Source:     s.AttrOr("data-source", ""),
		}
		locations = append(locations, location)
	})

	// 将结果编码为JSON格式并输出
	ToJSON(locations, "locationsdata.json", "./data/source/")
}

func GetTitlesFromHTML() {
	// Load the HTML file
	file, err := ioutil.ReadFile("./data/html/titles-list.html")
	if err != nil {
		panic(err)
	}

	// Parse the HTML with goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(file)))
	if err != nil {
		panic(err)
	}

	// Extract the div elements
	var titles []TitleSource
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		// Extract the text and class attributes of the div element
		text := s.Text()
		class, exists := s.Attr("class")
		if !exists {
			class = ""
		}

		// Append the title to the slice
		title := TitleSource{text, class}
		titles = append(titles, title)
	})

	// Encode the titles slice as JSON
	ToJSON(titles, "titlesdata.json", "./data/source/")
}
