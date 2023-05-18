package dao

import "time"

func CreateDeck(deck *Deck) error {
	return db.Create(&deck).Error
}

func GetDecksByCard(card Card) ([]Deck, error) {
	var decks []Deck
	err := db.Model(&card).Association("Decks").Find(&decks)
	return decks, err
}

func GetDecksByAuthorID(authorID uint) ([]Deck, error) {
	var decks []Deck
	err := db.Where("author_id = ?", authorID).Find(&decks).Error
	return decks, err
}

func DeleteDeckByID(deckID uint) error {
	// 删除关联关系
	if err := db.Model(&Deck{ID: deckID}).Association("Cards").Clear(); err != nil {
		return err
	}

	if err := db.Model(&Deck{ID: deckID}).Association("Author").Clear(); err != nil {
		return err
	}

	// 删除Deck
	if err := db.Delete(&Deck{ID: deckID}).Error; err != nil {
		return err
	}

	return nil
}

func GetRecentlyCreatedDecks(duration time.Duration) ([]Deck, error) {
	var decks []Deck
	now := time.Now()
	startTime := now.Add(-duration)

	err := db.Preload("Author").Preload("Cards").Where("created_at > ?", startTime).Order("created_at desc").Find(&decks).Error
	if err != nil {
		return nil, err
	}
	return decks, nil
}
