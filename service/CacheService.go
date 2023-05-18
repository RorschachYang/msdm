package service

import "os"

var CardsCache []Card
var LocationsCache []Location
var VariantsCache []Variant
var TitlesCache []Title
var TagsCache []Tag
var ArtistsCache []Artist
var appID string
var appSecret string
var uper string

func init() {
	CardsCache = GetAllCards()
	LocationsCache = GetAllLocations()
	VariantsCache = GetAllVariants()
	appID = os.Getenv("APP_ID")
	appSecret = os.Getenv("APP_SECRET")
	uper = os.Getenv("UPER")
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
