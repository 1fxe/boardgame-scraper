### Board Game Scraper & XML Client

Web scraper for [boardgamegeek.com](https://boardgamegeek.com/)

I also use the XML Api for the games see [xml_client.go](./xml_client.go)  

## Scraper Usage

```
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
  - [x] Name & Description
  - [ ] All Games
  - [x] Misc Data
  - [x] Age
  - [x] Parse Categories
  - [x] Parse Mechanics

- [x] Speed up parsing


### What is a Web Scraper
A web scraper is a tool or piece of software that is used to extract data from websites. Web scrapers are typically used to extract structured data from websites, such as contact information, prices, product descriptions, and other types of data that can be used for various purposes, such as market research or competitive analysis.

Web scrapers typically work by making HTTP requests to websites and then parsing the HTML or XML responses to extract the desired data. Some web scrapers are designed to mimic the actions of a human user, such as clicking on links and filling out forms, in order to access the data on websites that require user interaction. Other web scrapers are more focused on extracting data from websites that provide structured data, such as APIs or RSS feeds.

Overall, web scrapers are useful tools for extracting data from websites and can be used in a wide variety of applications.

### What is an XML Client
An XML API client is a piece of software that is specifically designed to access and interact with an application programming interface (API) that uses XML (Extensible Markup Language) as its data format. XML is a widely used data format that is used to structure and transmit data between different systems and applications.

An XML API client typically includes a set of programming libraries or classes that provide an interface for accessing an XML-based API. It also includes tools for making HTTP requests to the API and parsing the XML responses to extract the desired data.

XML API clients are commonly used by developers to access data and functionality from third-party APIs that use XML as their data format. For example, a developer might use an XML API client to access a financial data API that provides stock prices and other financial information in XML format.