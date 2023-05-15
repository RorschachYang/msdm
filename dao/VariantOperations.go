package dao

func CreateVariant(variant *Variant) error {
	err := db.Create(variant).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateVariant(variant *Variant) error {
	var existingVariant Variant
	err := db.Where("id = ?", variant.ID).First(&existingVariant).Error
	if err != nil {
		return err
	}

	existingVariant.Name = variant.Name
	existingVariant.Rarity = variant.Rarity
	existingVariant.CardID = variant.CardID
	existingVariant.ImageURLName = variant.ImageURLName
	existingVariant.Released = variant.Released

	return db.Save(&existingVariant).Error
}

func GetVariantByID(id uint) (*Variant, error) {
	var variant Variant
	result := db.First(&variant, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &variant, nil
}

func GetAllVariants() ([]Variant, error) {
	var variants []Variant
	result := db.Preload("Artist").Preload("Tags").Order("name asc").Find(&variants)
	if result.Error != nil {
		return nil, result.Error
	}
	return variants, nil
}
