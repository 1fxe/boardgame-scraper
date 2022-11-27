// scraper
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/1fxe/board-game-web-scraper/internal"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)

type TableConfig struct {
	CategoryURL     string `json:"categoriesURL"`
	MechanicsURL    string `json:"mechanicsURL"`
	Table           string `json:"table"`
	ItemName        string `json:"itemName"`
	ItemDescription string `json:"itemDescription"`
}

type Config struct {
	Domain      string      `json:"domain"`
	CreditsURL  string      `json:"creditsURL"`
	TableConfig TableConfig `json:"tableConfig"`
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
	var urls []URLforParsing

	s.collector.OnHTML(s.config.TableConfig.Table, func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			e.ForEach("td", func(_ int, e *colly.HTMLElement) {
				href := e.ChildAttr("a", "href")
				if href != "" {
					link := fmt.Sprintf("https://%s%s", s.config.Domain, href)
					urls = append(urls, URLforParsing{link})
				}
			})
		})
	})

	if table == internal.CategoriesURLS {
		_ = s.collector.Visit(s.config.TableConfig.CategoryURL)
	} else if table == internal.MechanicsURLS {
		_ = s.collector.Visit(s.config.TableConfig.MechanicsURL)
	}

	file, err := json.MarshalIndent(urls, "", "   ")
	if err != nil {
		log.Fatalln("Failed to marshal json", urls)
	}
	err = os.WriteFile(fmt.Sprintf("./data/%s.json", table), file, 0644)
	if err != nil {
		log.Fatalln("Failed to write json", file)
	}
}

func (s Scraper) parseFromTableURLs(table string) {
	file, err := os.ReadFile(fmt.Sprintf("./data/%sURLS.json", table))

	if err != nil {
		log.Printf("Error reading %s.json %s", table, err)
		return
	}

	var urlForParsing []URLforParsing

	err = json.Unmarshal(file, &urlForParsing)

	if err != nil {
		log.Printf("Error reading parsing json %s", err)
		return
	}

	log.Println("Started Parsing URLS")
	var data []internal.Data

	s.collector.OnXML(s.config.TableConfig.ItemName, func(e *colly.XMLElement) {
		s := strings.Split(e.Text, "GEEK.geekitemPreload = ")[1]
		s = strings.Split(s, "GEEK.geekitemSettings = ")[0]
		s = strings.TrimSpace(s)
		s = strings.TrimSuffix(s, ";")
		var autoGen internal.GeekScriptItem
		_ = json.Unmarshal([]byte(s), &autoGen)

		log.Println(autoGen.Item.Name)

		// TODO Improve this
		cleanDescription := strings.Split(autoGen.Item.Description, "/em> ")
		if len(cleanDescription) > 1 {
			data = append(data, internal.Data{Name: autoGen.Item.Name, Description: cleanDescription[1]})
		} else {
			data = append(data, internal.Data{Name: autoGen.Item.Name, Description: cleanDescription[0]})
		}
	})

	for _, url := range urlForParsing {
		_ = s.collector.Visit(url.URL)
	}

	log.Println("Finished Parsing URLS")

	out, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		log.Fatalln("Failed to marshal json", data)
	}
	err = os.WriteFile(fmt.Sprintf("./data/%s.json", table), out, 0644)
	if err != nil {
		log.Fatalln("Failed to write json", out)
	}
}

func main() {

	getCategories := flag.Bool("getCategories", false, "Gets list of categories for future parsing")
	parseCategories := flag.Bool("parseCategories", false, "Parse categories from list")

	getMechanics := flag.Bool("getMechanics", false, "Gets list of mechanism for future parsing")
	parseMechanics := flag.Bool("parseMechanics", false, "Parse mechanisms from list")

	flag.Parse()

	cfgFile, err := os.ReadFile("config.json")

	if err != nil {
		log.Printf("Error reading config.json %s", err)
		return
	}

	cfg := Config{}

	err = json.Unmarshal(cfgFile, &cfg)

	if err != nil {
		log.Printf("Error parsing json %s", err)
		return
	}

	log.Println("Successfully loaded config:", cfg)

	if err := os.Mkdir("data", os.ModePerm); err != nil {
		log.Println(err)
	}

	options := append(chromedp.DefaultExecAllocatorOptions[:],
		// block all images
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
	)
	allocatorCtx, cancel := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocatorCtx)
	defer cancel()

	collector := colly.NewCollector(
		colly.AllowedDomains(cfg.Domain),
	)

	collector.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	collector.OnResponse(func(r *colly.Response) {
		log.Println("Visited", r.Request.URL)
	})

	scraper := Scraper{cfg, ctx, collector}

	if *getCategories {
		scraper.getDataFromTable(internal.CategoriesURLS)
	}

	if *parseCategories {
		scraper.parseFromTableURLs(internal.Categories)
	}

	if *getMechanics {
		scraper.getDataFromTable(internal.MechanicsURLS)
	}

	if *parseMechanics {
		scraper.parseFromTableURLs(internal.Mechanics)
	}

}
