// xml_client
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/1fxe/board-game-web-scraper/internal"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// boardgamegeek api
// https://boardgamegeek.com/wiki/page/BGG_XML_API2

const (
	API  = "https://www.boardgamegeek.com/xmlapi2/thing"
	Type = "boardgame"
)

func profile(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func loadBoardGameProperties(s string) map[string]internal.Data {
	file, err := os.ReadFile(fmt.Sprintf("./data/%s.json", internal.Categories))

	if err != nil {
		log.Panicln("Error reading file: ", err)
	}

	var data []internal.Data

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Panicln("Error unmarshalling file: ", err)
	}

	var dataMap = make(map[string]internal.Data)
	for _, datum := range data {
		dataMap[datum.Name] = datum
	}

	return dataMap
}

func main() {
	var allCategories = loadBoardGameProperties(internal.Categories)
	var allMechanics = loadBoardGameProperties(internal.Mechanics)

	// fetch data from API
	var listOfIds []string
	for i := 1; i < 100; i++ {
		listOfIds = append(listOfIds, strconv.Itoa(i))
	}

	start := time.Now()
	resp, err := http.Get(API + fmt.Sprintf("?id=%s&type=%s", strings.Join(listOfIds, ","), Type))
	if err != nil {
		log.Fatalln("Failed to fetch data from API", err)
	}
	profile(start, "fetch data from API")

	if resp.StatusCode != 200 {
		log.Fatalln("Failed to fetch data from API", resp.StatusCode)
	}

	start = time.Now()
	var data internal.Items
	err = xml.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalln("Failed to unmarshal data", err)
	}
	profile(start, "unmarshal data")

	var games []internal.BoardGame

	start = time.Now()
	for _, item := range data.Item {
		var categories []internal.Data
		var mechanisms []internal.Data
		for _, link := range item.Link {
			if link.Type == internal.Categories {
				category := allCategories[link.Value]
				if category != (internal.Data{}) {
					categories = append(categories, category)
				}
			} else if link.Type == internal.Mechanics {
				mechanism := allMechanics[link.Value]
				if mechanism != (internal.Data{}) {
					mechanisms = append(mechanisms, mechanism)
				}
			}
		}
		games = append(games, internal.BoardGame{
			Name:         item.Name[0].Value,
			Description:  item.Description,
			YearReleased: item.Yearpublished.Value,
			NoPlayers: internal.Range{
				Min: item.Minplayers.Value,
				Max: item.Maxplayers.Value,
			},
			PlayTime: internal.Range{
				Min: item.Minplaytime.Value,
				Max: item.Maxplaytime.Value,
			},
			MinAge: item.Minage.Value,
			Characteristic: internal.Characteristic{
				Categories: categories,
				Mechanisms: mechanisms,
			},
		})
	}

	out, err := json.MarshalIndent(games, "", "   ")
	if err != nil {
		log.Fatalln("Failed to marshal json", data)
	}
	err = os.WriteFile(fmt.Sprintf("./data/games_%d.json", 1), out, 0644)
	if err != nil {
		log.Fatalln("Failed to write json", out)
	}
	profile(start, "marshal data")
}
