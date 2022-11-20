// scraper
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/1fxe/board-game-web-scraper/internal"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
)

type CategoriesConfig struct {
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
}

type TableConfig struct {
	CategoryURL  string `json:"categoriesURL"`
	MechanicsURL string `json:"mechanicsURL"`
	Table        string `json:"table"`
}

type Config struct {
	Domain           string           `json:"domain"`
	CreditsURL       string           `json:"creditsURL"`
	TableConfig      TableConfig      `json:"tableConfig"`
	CategoriesConfig CategoriesConfig `json:"categories"`
}

type URLforParsing struct {
	URL string `json:"url"`
}

type Scraper struct {
	config    Config
	ctx       context.Context
	collector *colly.Collector
}

func (s Scraper) getDataFromTable(table string) {
	urls := []URLforParsing{}

	s.collector.OnHTML(s.config.TableConfig.Table, func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			e.ForEach("td", func(_ int, e *colly.HTMLElement) {
				link := e.ChildAttr("a", "href")
				urls = append(urls, URLforParsing{link})
			})
		})
	})

	if table == internal.Categories {
		s.collector.Visit(s.config.TableConfig.CategoryURL)
	} else if table == internal.Mechanics {
		s.collector.Visit(s.config.TableConfig.MechanicsURL)
	}

	catsFile, _ := json.MarshalIndent(urls, "", "   ")
	_ = ioutil.WriteFile(fmt.Sprintf("./data/%s.json", table), catsFile, 0644)
}

func (s Scraper) parseCategories() {
	categoriesFile, err := ioutil.ReadFile(fmt.Sprintf("./data/%s.json", internal.Categories))

	if err != nil {
		fmt.Printf("Error reading ceategories.json %s", err)
		return
	}

	urlForParsing := []URLforParsing{}

	err = json.Unmarshal(categoriesFile, &urlForParsing)

	if err != nil {
		fmt.Printf("Error reading parsing json %s", err)
		return
	}

	categories := []internal.Category{}

	for _, category := range urlForParsing {

		url := fmt.Sprintf("https://%s%s", s.config.Domain, category.URL)
		var categoryName, categoryDescription string

		if err := chromedp.Run(s.ctx,
			chromedp.Navigate(url),
			chromedp.Text(s.config.CategoriesConfig.CategoryName, &categoryName, chromedp.ByQuery),
			chromedp.Text(s.config.CategoriesConfig.Description, &categoryDescription, chromedp.ByQuery),
		); err != nil {
			fmt.Println(err)
		}

		cleanDescription := strings.Split(categoryDescription, "\n\nMicrobadges\n\n")[0]
		categories = append(categories, internal.Category{categoryName, cleanDescription})

	}

	catsFile, _ := json.MarshalIndent(categories, "", "   ")
	_ = ioutil.WriteFile("./data/categories.json", catsFile, 0644)
}

func main() {

	verbose := flag.Bool("verbose", false, "Prints a lot of information")
	getCategories := flag.Bool("getCategories", false, "Gets list of categories for future parsing")
	parseCategories := flag.Bool("parseCategories", false, "Parse categories from list")

	getMechanics := flag.Bool("getMechanics", false, "Gets list of mechanism for future parsing")
	parseMechanics := flag.Bool("parseMechanis", false, "Parse mechanisms from list")

	flag.Parse()

	cfgFile, err := ioutil.ReadFile("config.json")

	if err != nil {
		fmt.Printf("Error reading config.json %s", err)
		return
	}

	cfg := Config{}

	err = json.Unmarshal(cfgFile, &cfg)

	if err != nil {
		fmt.Printf("Error parsing json %s", err)
		return
	}

	if *verbose {
		fmt.Println("Succefully loaded config:", cfg)
	}

	if err := os.Mkdir("data", os.ModePerm); err != nil {
		if *verbose {
			fmt.Println(err)
		}
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	collector := colly.NewCollector(
		colly.AllowedDomains(cfg.Domain),
	)

	if *verbose {
		collector.OnError(func(_ *colly.Response, err error) {
			fmt.Println("Something went wrong:", err)
		})

		collector.OnResponse(func(r *colly.Response) {
			fmt.Println("Visited", r.Request.URL)
		})
	}

	scraper := Scraper{cfg, ctx, collector}

	if *getCategories {
		scraper.getDataFromTable(internal.Categories)
	}

	if *parseCategories {
		scraper.parseCategories()
	}

	if *getMechanics {
		scraper.getDataFromTable(internal.Mechanics)
	}

	if *parseMechanics {
		scraper.parseCategories()
	}

}
