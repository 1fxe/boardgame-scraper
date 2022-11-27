### Board Game Scraper

Web scraper for [boardgamegeek.com](https://boardgamegeek.com/)

I also use the XML Api for the games see [xml_client.go](./xml_client.go)  

## Scraper Usage

```go
-getCategories
    Gets list of categories for future parsing
-getMechanics
    Gets list of mechanism for future parsing
-parseCategories
    Parse categories from list
-parseMechanics
    Parse mechanisms from list
```

## TODO

- [x] Parse Mechanics
- [x] Parse Categories
- [ ] Parse Game Data
  - [ ] Name & Description
  - [ ] All Games
  - [x] Misc Data
  - [x] Age
  - [ ] Parse Categories
  - [ ] Parse Mechanics

- [x] Speed up parsing
