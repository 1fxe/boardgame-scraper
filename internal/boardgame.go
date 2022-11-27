// Package internal
package internal

const (
	CategoriesURLS = "categoriesURLS"
	MechanicsURLS  = "mechanicsURLS"
	Categories     = "categories"
	Mechanics      = "mechanics"
)

type Data struct {
	Name        string `json:"dataName"`
	Description string `json:"dataDescription"`
}

type Range struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Game struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	YearReleased int    `json:"yearReleased"`
	NoPlayers    Range  `json:"noPlayers"`
	PlayTime     Range  `json:"playTime"`
	MinAge       int    `json:"age"`
	Categories   []Data `json:"categories"`
	Mechanisms   []Data `json:"mechanisms"`
}

type GeekScriptItem struct {
	Item struct {
		Href        string `json:"href,omitempty"`
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"item,omitempty"`
}

type Items struct {
	Item []struct {
		Text      string `xml:",chardata"`
		Type      string `xml:"type,attr"`
		ID        string `xml:"id,attr"`
		Thumbnail string `xml:"thumbnail"`
		Image     string `xml:"image"`
		Name      []struct {
			Text      string `xml:",chardata"`
			Type      string `xml:"type,attr"`
			Sortindex string `xml:"sortindex,attr"`
			Value     string `xml:"value,attr"`
		} `xml:"name"`
		Description   string `xml:"description"`
		Yearpublished struct {
			Text  string `xml:",chardata"`
			Value int    `xml:"value,attr"`
		} `xml:"yearpublished"`
		Minplayers struct {
			Text  string `xml:",chardata"`
			Value int    `xml:"value,attr"`
		} `xml:"minplayers"`

		Playingtime struct {
			Text  string `xml:",chardata"`
			Value string `xml:"value,attr"`
		} `xml:"playingtime"`
		Minplaytime struct {
			Text  string `xml:",chardata"`
			Value int    `xml:"value,attr"`
		} `xml:"minplaytime"`
		Maxplaytime struct {
			Text  string `xml:",chardata"`
			Value int    `xml:"value,attr"`
		} `xml:"maxplaytime"`
		Minage struct {
			Text  string `xml:",chardata"`
			Value int    `xml:"value,attr"`
		} `xml:"minage"`
		Maxplayers struct {
			Text  string `xml:",chardata"`
			Value int    `xml:"value,attr"`
		} `xml:"maxplayers"`
	} `xml:"item"`
}
