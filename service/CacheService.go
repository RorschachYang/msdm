package service

var CardsCache []Card
var LocationsCache []Location
var VariantsCache []Variant
var TitlesCache []Title
var TagsCache []Tag
var ArtistsCache []Artist

func init() {
	CardsCache = GetAllCards()
	LocationsCache = GetAllLocations()
	VariantsCache = GetAllVariants()
}

func GetVariantsCache() []Variant {
	return VariantsCache
}

func GetCardsCache() []Card {
	return CardsCache
}

func GetLocationsCache() []Location {
	return LocationsCache
}
