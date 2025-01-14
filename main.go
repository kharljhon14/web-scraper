package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Scraping logic

	c := colly.NewCollector(colly.AllowedDomains("www.scrapingcourse.com"))

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Printf("%v", e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scrapped!")
	})

	c.Visit("https://www.scrapingcourse.com/ecommerce")
}
