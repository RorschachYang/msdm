package dao

func CreateTitle(title *Title) error {
	return db.Create(title).Error
}

func UpdateTitle(title *Title) error {
	var existingTitle Title
	err := db.Where("name = ?", title.Name).First(&existingTitle).Error
	if err != nil {
		return err
	}

	existingTitle.NameZh = title.NameZh
	return db.Save(&existingTitle).Error
}

func GetTitleByName(name string) (*Title, error) {
	var title Title
	result := db.Where("name = ?", name).First(&title)
	if result.Error != nil {
		return nil, result.Error
	}
	return &title, nil
}

func GetAllTitles() ([]Title, error) {
	var titles []Title
	result := db.Find(&titles)
	if result.Error != nil {
		return nil, result.Error
	}
	return titles, nil
}
