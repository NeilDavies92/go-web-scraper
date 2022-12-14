package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json: "name"`
	Price  string `json: "price"`
	ImgUrl string `json: "imageurl"`
}

func main() {

	// Colly’s main entity is a Collector object
	// Collector manages the network communication and responsible for the
	// execution of the attached callbacks while a collector job is running.

	c := colly.NewCollector(
		colly.AllowedDomains("j2store.net"),
	)

	var items []item

	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		item := item{
			Name:   h.ChildText("h2.product-title"),
			Price:  h.ChildText("div.sale-price"),
			ImgUrl: h.ChildAttr("img", "src"),
		}

		items = append(items, item)
	})

	// Find and visit next page
	c.OnHTML("[title=Next]", func(h *colly.HTMLElement) {
		nextPage := h.Request.AbsoluteURL(h.Attr("href"))
		c.Visit(nextPage)
	})

	// Print out URL of pages scraped
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	c.Visit("https://j2store.net/demo/index.php/shop")

	// Save scraped contents as json file
	content, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err)
	}

	os.WriteFile("items.json", content, 0644)
}
