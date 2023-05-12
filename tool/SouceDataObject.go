package tool

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
