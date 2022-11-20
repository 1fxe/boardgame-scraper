// scraper
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type Categories struct {
	Table string `json: "table"`
}

type Config struct {
	Domain        string     `json: "domain"`
	CreditsURL    string     `json: "creditsURL"`
	CategoriesURL string     `json: "categoriesURL"`
	Categories    Categories `json: "categories"`
}

type CategoryForParsing struct {
	Name string `json: "categoryName"`
	URL  string `json: categoryURL`
}

func main() {
	cfgFile, err := ioutil.ReadFile("config.json")

	if err != nil {
		fmt.Printf("Error reading config.json %s", err)
		return
	}

	cfg := Config{}

	err = json.Unmarshal(cfgFile, &cfg)

	if err != nil {
		fmt.Printf("Error reading parsin json %s", err)
		return
	}

	fmt.Println(cfg)

	c := colly.NewCollector(
		colly.AllowedDomains(cfg.Domain),
	)

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	catforParsing := []CategoryForParsing{}

	c.OnHTML(cfg.Categories.Table, func(e *colly.HTMLElement) {
		// Each row has two columns
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			// Each column has an "a" tag
			e.ForEach("td", func(_ int, e *colly.HTMLElement) {
				// Each tag contains the category name and link in href
				e.ForEach("a", func(_ int, e *colly.HTMLElement) {
					name := e.Text
					link := e.Attr("href")
					catforParsing = append(catforParsing, CategoryForParsing{name, link})
					// fmt.Printf("Name: %s, link %s\n", name, link)
				})
			})
		})
	})

	c.Visit(cfg.CategoriesURL)

	catsFile, _ := json.MarshalIndent(catforParsing, "", "   ")

	_ = ioutil.WriteFile("catsForParsing.json", catsFile, 0644)
}
