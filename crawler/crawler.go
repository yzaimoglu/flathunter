package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
  c := colly.NewCollector()
 
  // Find and visit all links
	c.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.Visit("http://go-colly.org/")
}
