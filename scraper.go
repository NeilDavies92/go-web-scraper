package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Item struct {
	Name     string `json: "name"`
	Price    string `json: "price"`
	ImageUrl string `json: "imageurl"`
}

func main() {
	// Colly’s main entity is a Collector object
	// Collector manages the network communication and responsible for the
	// execution of the attached callbacks while a collector job is running.

	domain := "j2store.net"

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
	)

	c.OnHTML("div.col-sm-9 div[itemprop=itemListElement]", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildText("h2.product-title"))
	})

	url := "https://j2store.net/demo/index.php/shop"

	c.Visit(url)
}
