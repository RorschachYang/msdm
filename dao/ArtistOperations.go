package dao

func UpsertArtistAndLinkWithVariant(name string, vid uint) error {
	var artist Artist
	err := db.Where("name = ?", name).FirstOrCreate(&artist, Artist{Name: name}).Error
	if err != nil {
		return err
	}

	var variant Variant
	err = db.First(&variant, vid).Error
	if err != nil {
		return err
	}

	err = db.Model(&variant).Association("Artist").Append(&artist)
	if err != nil {
		return err
	}

	return nil
}

func GetAllArtists() ([]Artist, error) {
	var artists []Artist
	result := db.Find(&artists)
	if result.Error != nil {
		return nil, result.Error
	}
	return artists, nil
}
