package dao

func CreateLocation(location *Location) error {
	err := db.Create(location).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateLocation(location *Location) error {
	var existingLocation Location
	err := db.Where("name = ?", location.Name).First(&existingLocation).Error
	if err != nil {
		return err
	}

	existingLocation.NameZh = location.NameZh
	existingLocation.Description = location.Description
	existingLocation.DescriptionZh = location.DescriptionZh
	existingLocation.ImageURLName = location.ImageURLName
	existingLocation.Released = location.Released
	existingLocation.DefID = location.DefID

	err = db.Save(&existingLocation).Error
	if err != nil {
		return err
	}

	return nil
}

func GetLocationByName(name string) (*Location, error) {
	var location Location
	result := db.Where("name = ?", name).First(&location)
	if result.Error != nil {
		return nil, result.Error
	}
	return &location, nil
}

func GetAllLocations() ([]Location, error) {
	var locations []Location
	result := db.Order("name asc").Find(&locations)
	if result.Error != nil {
		return nil, result.Error
	}
	return locations, nil
}
