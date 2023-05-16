package service

import (
	"fmt"
	"regexp"

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

	fmt.Println(cardNames)
	fmt.Println(name)
	fmt.Println(description)
	fmt.Println(openid)
	cards, _ := dao.GetCardsByNames(cardNames)
	for _, card := range cards {
		fmt.Println(card.Name)
		fmt.Println(card.NameZh)
	}
	// user, _ := dao.GetUserByOpenID(openid)

	// newDeck := &dao.Deck{
	// 	Name:        name,
	// 	Description: description,
	// 	Code:        code,
	// 	Cards:       cards,
	// 	Author:      *user,
	// }

	// dao.CreateDeck(newDeck)
}
