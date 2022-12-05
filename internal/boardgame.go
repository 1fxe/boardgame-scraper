// Package internal
package internal

const (
	Categories     = "boardgamecategory"
	Mechanics      = "boardgamemechanic"
	CategoriesURLS = Categories + "URLS"
	MechanicsURLS  = Mechanics + "URLS"
)

type Characteristic struct {
	Categories []Data `json:"categories,omitempty"`
	Mechanisms []Data `json:"mechanisms,omitempty"`
}

type Data struct {
	Name        string `json:"dataName"`
	Description string `json:"dataDescription"`
}

type Range struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type BoardGame struct {
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	YearReleased   int            `json:"yearReleased"`
	NoPlayers      Range          `json:"noPlayers"`
	PlayTime       Range          `json:"playTime"`
	MinAge         int            `json:"age"`
	Characteristic Characteristic `json:"characteristic,omitempty"`
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
			Value string `xml:"value,attr"`
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
			Value int    `xml:"value,attr"`
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
		Link []struct {
			Type  string `xml:"type,attr"`
			Value string `xml:"value,attr"`
		} `xml:"link"`
	} `xml:"item"`
}
