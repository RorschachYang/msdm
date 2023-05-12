package service

import "github.com/RorschachYang/msdm/dao"

func GetAllLocations() []Location {
	locationsDO, err := dao.GetAllLocations()
	if err == nil {
		//中文翻译不在这里更新
		var locations []Location
		for _, locationsDO := range locationsDO {
			location := Location{
				Name:               locationsDO.Name,
				NameZh:             locationsDO.NameZh,
				Description:        locationsDO.Description,
				DescriptionZh:      locationsDO.DescriptionZh,
				DefID:              locationsDO.DefID,
				ImageURL:           ImageURLPrefix + locationsDO.ImageURLName + ".webp",
				ImageURLCompressed: CompressedImageURLPrefix + "locations/" + locationsDO.ImageURLName + ".webp",
				Released:           locationsDO.Released,
			}
			locations = append(locations, location)
		}
		return locations
	}
	return nil
}
