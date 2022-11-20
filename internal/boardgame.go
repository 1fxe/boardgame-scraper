// boardgame
package internal

const (
	Categories = "categoriesURLS"
	Mechanics  = "mechanicsURLS"
)

type Category struct {
	Name        string `json:"categoryName"`
	Description string `json:"categoryDescription"`
}

type Mechanism struct {
	Name        string `json:"mechanismName"`
	Description string `json:"mechanismDescription"`
}

type Range struct {
	Min uint16 `json:"min"`
	Max uint16 `json:"max"`
}

type Game struct {
	PrimaryName  string      `json:"name"`
	YearReleased uint16      `json:"yearReleased"`
	NoPlayers    Range       `json:"noPlayers"`
	PlayTime     Range       `json:"playTime"`
	Age          string      `json:"age"`
	Categories   []Category  `json:"categories"`
	Mechanisms   []Mechanism `json:"mechanisms"`
}
