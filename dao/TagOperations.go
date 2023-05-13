package dao

func UpsertTag(tag *Tag) error {
	var existingTag Tag
	err := db.Where("name = ?", tag.Name).FirstOrCreate(&existingTag, tag).Error
	if err != nil {
		return err
	}

	// Update existing tag
	existingTag.NameZh = tag.NameZh
	return db.Save(&existingTag).Error
}

func UpsertTagAndLinkWithVariant(tagName string, variantID uint) error {
	// 检查tag是否存在，若不存在则创建
	var tag Tag
	err := db.Where("name = ?", tagName).FirstOrCreate(&tag, Tag{Name: tagName}).Error
	if err != nil {
		return err
	}

	// 检查variant是否存在
	var variant Variant
	err = db.Where("id = ?", variantID).First(&variant).Error
	if err != nil {
		return err
	}

	// 将tag关联到variant
	err = db.Model(&variant).Association("Tags").Append(&tag)
	if err != nil {
		return err
	}

	return nil
}

func GetAllTags() ([]Tag, error) {
	var tags []Tag
	result := db.Order("name asc").Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}
	return tags, nil
}
