package dao

func CreateCard(card *Card) error {
	return db.Create(card).Error
}

func UpdateCard(card *Card) error {
	var existingCard Card
	err := db.Where("id = ?", card.ID).First(&existingCard).Error
	if err != nil {
		return err
	}

	existingCard.Name = card.Name
	existingCard.NameZh = card.NameZh
	existingCard.Description = card.Description
	existingCard.DescriptionZh = card.DescriptionZh
	existingCard.Source = card.Source
	existingCard.Power = card.Power
	existingCard.Cost = card.Cost
	existingCard.ImageURLName = card.ImageURLName
	existingCard.Defid = card.Defid
	return db.Save(&existingCard).Error
}

func GetCardByID(id uint) (*Card, error) {
	var card Card
	result := db.First(&card, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &card, nil
}

func GetAllCards() ([]Card, error) {
	var cards []Card
	result := db.Order("name asc").Find(&cards)
	if result.Error != nil {
		return nil, result.Error
	}
	return cards, nil
}
