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

func main() {
	// fetch data from API
	var listOfIds []string
	for i := 1; i < 1_000; i++ {
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

	var games []internal.Game

	start = time.Now()
	for _, item := range data.Item {
		games = append(games, internal.Game{
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
			MinAge:     item.Minage.Value,
			Categories: []internal.Data{},
			Mechanisms: []internal.Data{},
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
