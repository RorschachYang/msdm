package main

import (
	"fmt"
	"net/http"

	"github.com/RorschachYang/msdm/controller"
	"github.com/RorschachYang/msdm/tool"
)

func main() {

	tool.CardsTranslationToJson()
	tool.LocationTranslationToJson()
	tool.TagTranslationToJson()

	http.HandleFunc("/listCards", controller.ListCards)
	http.HandleFunc("/listVariants", controller.ListVariants)
	http.HandleFunc("/listLocations", controller.ListLocations)
	http.HandleFunc("/getCard", controller.GetCard)
	http.HandleFunc("/getVariant", controller.GetVariant)
	http.HandleFunc("/getVariantsByCid", controller.GetVariantsByCardID)
	http.HandleFunc("/getLocation", controller.GetLocation)
	http.HandleFunc("/login", controller.Login)

	fmt.Println("Server started at port:80")
	http.ListenAndServe(":80", nil)
}
