package dao

import "gorm.io/gorm"

func CreateDeck(deck *Deck) error {
	return db.Create(&deck).Error
}

func GetAllDecks(db *gorm.DB) ([]Deck, error) {
	var decks []Deck
	err := db.Find(&decks).Error
	return decks, err
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
	deck := Deck{ID: deckID}
	return db.Delete(&deck).Error
}
