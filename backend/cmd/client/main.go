package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/yzaimoglu/flathunter/pkg/crawler"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

func main() {
	// Declaring alternating user agents
	round_robin_ua, _ := crawler.NewUserAgent(
		&models.UserAgent{ID: 1, UserAgent: "Mozilla/5.0 (Windows NT 5.1; en-US; rv:1.9.1.20) Gecko/20140810 Firefox/37.0"},
		&models.UserAgent{ID: 2, UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84"},
		&models.UserAgent{ID: 3, UserAgent: "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_5_2 rv:6.0) Gecko/20130127 Firefox/36.0"},
		&models.UserAgent{ID: 4, UserAgent: "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_6_5 rv:6.0; sl-SI) AppleWebKit/531.49.6 (KHTML, like Gecko) Version/5.0.1 Safari/531.49.6"},
	)

	// round_robin_proxy, _ := utilities.NewProxy(
	// )

	// Initializing the listings slice and the colly collector
	var listings []models.Listing = []models.Listing{}
	c := colly.NewCollector(
		colly.UserAgent("Flathunters"),
		colly.AllowURLRevisit(),
		colly.CacheDir("./cache"),
		colly.MaxDepth(2),
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4})
	c.SetRequestTimeout(120 * time.Second)

	// Setting proxy
	//c.SetProxy(utilities.ProxyString(round_robin_proxy.Next()))

	// Cloning the colly collector for the detailCollector
	detailCollector := c.Clone()

	// Setting the alternating User Agent
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", round_robin_ua.Next().UserAgent)
	})
	detailCollector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", round_robin_ua.Next().UserAgent)
	})

	// Visiting the listings specific urls to scrape
	c.OnHTML("article.aditem", func(e *colly.HTMLElement) {
		detailCollector.Visit("https://ebay-kleinanzeigen.de" + e.ChildAttr("a.ellipsis", "href"))
	})

	// Error while scraping
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	// Scraping the listings
	detailCollector.OnHTML("article[id=viewad-product]", func(e *colly.HTMLElement) {
		fmt.Println(e.Request.Headers.Get("User-Agent"))
		// Setting initial settings for Ebay-Kleinanzeigen
		var listing models.Listing = models.Listing{
			URLID: 1,
			URL:   e.Request.URL.Host,
		}
		// Scraping price
		listing.Price = e.ChildText("h2[id=viewad-price]")
		images := []string{}

		// Scraping date
		date := e.ChildText("div[id=viewad-extra-info] span")
		listing.Date = date

		// Scraping description
		description := e.ChildText("p[id=viewad-description-text]")
		listing.Description = description

		// Scraping details
		details := []string{}
		e.ForEach("div[id=viewad-details] li", func(_ int, detail *colly.HTMLElement) {
			details = append(details, detail.Text)
		})
		listing = crawler.GetDetailsEbay(details, listing)

		// Scraping images
		e.ForEach("img[id=viewad-image]", func(_ int, image *colly.HTMLElement) {
			images = append(images, image.Attr("src"))
		})
		listing.Images = images
		listings = append(listings, listing)
	})

	// Visiting and waiting
	c.Visit("https://www.ebay-kleinanzeigen.de/s-wandsbek/wohnung/k0l9446")
	c.Wait()

	// Waiting for the jobs to be complete
	time.Sleep(5 * time.Second)

	// Looping through the listings and printing out the resulting objects
	for i := range listings {
		fmt.Println(listings[i])
	}
}
