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
