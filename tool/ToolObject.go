package tool

type Card struct {
	Name               string    `json:"name"`
	NameZh             string    `json:"nameZh"`
	Description        string    `json:"description"`
	DescriptionZh      string    `json:"descriptionZh"`
	Source             string    `json:"source"`
	Power              int       `json:"power"`
	Cost               int       `json:"cost"`
	ImageURL           string    `json:"imageURL"`
	ImageURLCompressed string    `json:"imageURLCompressed"`
	Tags               []string  `json:"tags"`
	Cid                string    `json:"cid"`
	Defid              string    `json:"defid"`
	Variants           []Variant `json:"variants,omitempty"`
}

type Variant struct {
	Name               string   `json:"name"`
	NameZh             string   `json:"nameZh"`
	ImageURL           string   `json:"imageURL"`
	FullImagSrc        string   `json:"fullImgSrc,omitempty"`
	Artist             string   `json:"artist"`
	Inker              string   `json:"inker"`
	Colorist           string   `json:"colorist"`
	Tags               []string `json:"tags"`
	Rarity             string   `json:"rarity"`
	Cid                string   `json:"cid"`
	Vid                string   `json:"vid"`
	Released           bool     `json:"released"`
	ImageURLCompressed string   `json:"imageURLCompressed"`
}

type Artist struct {
	Name string `json:"name"`
}

type Location struct {
	Name               string `json:"name"`
	NameZh             string `json:"nameZh"`
	Description        string `json:"description"`
	DescriptionZh      string `json:"descriptionZh"`
	ImageURL           string `json:"imageURL"`
	DefID              string `json:"defId"`
	Released           bool   `json:"released"`
	ImageURLCompressed string `json:"imageURLCompressed"`
}

type Title struct {
	Name   string `json:"name"`
	NameZh string `json:"nameZh"`
}

type VariantTag struct {
	Value string `json:"value"`
}

type CardSource struct {
	Name        string   `json:"name"`
	Cost        string   `json:"cost"`
	Power       string   `json:"power"`
	Ability     string   `json:"ability"`
	Src         string   `json:"src"`
	CID         string   `json:"cid"`
	Rarity      string   `json:"rarity"`
	Sketcher    string   `json:"sketcher"`
	Colorist    string   `json:"colorist"`
	Inker       string   `json:"inker"`
	Status      string   `json:"status"`
	CardDefID   string   `json:"carddefid"`
	Source      string   `json:"source"`
	Tags        string   `json:"tags"`
	Variants    []string `json:"variants"`
	Description string   `json:"description"`
	HRef        string   `json:"href"`
	CardName    string   `json:"cardname"`
}

type VariantSource struct {
	Name        string `json:"name"`
	Cost        string `json:"cost"`
	Power       string `json:"power"`
	Ability     string `json:"ability"`
	Src         string `json:"src"`
	Cid         string `json:"cid"`
	Vid         string `json:"vid"`
	Rarity      string `json:"rarity"`
	Sketcher    string `json:"sketcher"`
	Colorist    string `json:"colorist"`
	Inker       string `json:"inker"`
	Status      string `json:"status"`
	CardDefID   string `json:"carddefid"`
	Source      string `json:"source"`
	Tags        string `json:"tags"`
	Description string `json:"description"`
	HRef        string `json:"href"`
}

type TitleSource struct {
	Text  string `json:"text"`
	Class string `json:"class"`
}

type ArtistSource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationSource struct {
	Name       string `json:"name"`
	Cost       string `json:"cost"`
	Power      string `json:"power"`
	Ability    string `json:"ability"`
	ImageSrc   string `json:"image_src"`
	Rarity     string `json:"rarity"`
	Difficulty string `json:"difficulty"`
	Status     string `json:"status"`
	CardDefID  string `json:"card_def_id"`
	Source     string `json:"source"`
}

type CardTranslation struct {
	Cid           string `json:"cid"`
	Description   string `json:"description"`
	DescriptionZh string `json:"descriptionZh"`
	NameZh        string `json:"nameZh"`
	Name          string `json:"name"`
}

type LocationTranslation struct {
	Name          string `json:"name"`
	NameZh        string `json:"nameZh"`
	Description   string `json:"description"`
	DescriptionZh string `json:"descriptionZh"`
	DefID         string `json:"defId"`
}

type TagTranslation struct {
	Name   string `json:"name"`
	NameZh string `json:"nameZh"`
}
