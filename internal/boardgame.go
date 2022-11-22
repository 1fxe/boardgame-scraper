// boardgame
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
	Min uint16 `json:"min"`
	Max uint16 `json:"max"`
}

type Game struct {
	PrimaryName  string `json:"name"`
	YearReleased uint16 `json:"yearReleased"`
	NoPlayers    Range  `json:"noPlayers"`
	PlayTime     Range  `json:"playTime"`
	Age          string `json:"age"`
	Categories   []Data `json:"categories"`
	Mechanisms   []Data `json:"mechanisms"`
}
