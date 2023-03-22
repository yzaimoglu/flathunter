package crawler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

// GetDetailsEbay scrapes the details of a listing from the ebay-kleinanzeigen website
func GetDetailsEbay(details []string, listing models.Listing) (resultingListing models.Listing) {
	// Loop through scraped details and harvest specific details
	for i := range details {
		details[i] = strings.ReplaceAll(strings.ReplaceAll(details[i], " ", ""), "\n", "")
		replacer := ""
		if strings.HasPrefix(details[i], "Wohnfläche") {
			replacer = "Wohnfläche"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			listing.Size = details[i]
		} else if strings.HasPrefix(details[i], "Zimmer") {
			replacer = "Zimmer"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			rooms_int, err := strconv.Atoi(details[i])
			if err != nil {
				fmt.Println("Error during conversion of rooms to int")
				rooms_int = 0
			}
			listing.Rooms = rooms_int
		} else if strings.HasPrefix(details[i], "Badezimmer") {
			replacer = "Badezimmer"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			bathrooms_int, err := strconv.Atoi(details[i])
			if err != nil {
				fmt.Println("Error during conversion of bathrooms to int")
				bathrooms_int = 0
			}
			listing.Bathrooms = bathrooms_int
		} else if strings.HasPrefix(details[i], "Etage") {
			replacer = "Etage"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			listing.Floor = details[i]
		} else if strings.HasPrefix(details[i], "Wohnungstyp") {
			replacer = "Wohnungstyp"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			listing.Type = details[i]
		} else if strings.HasPrefix(details[i], "Nebenkosten") {
			replacer = "Nebenkosten"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			listing.ExtraCosts = details[i]
		} else if strings.HasPrefix(details[i], "Warmmiete") {
			replacer = "Warmmiete"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			listing.FullRent = details[i]
		} else if strings.HasPrefix(details[i], "Kaution/Genoss.-Anteile") {
			replacer = "Kaution/Genoss.-Anteile"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			listing.Deposit = details[i]
		} else if strings.HasPrefix(details[i], "Schlafzimmer") {
			replacer = "Schlafzimmer"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			bedrooms_int, err := strconv.Atoi(details[i])
			if err != nil {
				fmt.Println("Error during conversion of bedrooms to int")
				bedrooms_int = 0
			}
			listing.Bedrooms = bedrooms_int
		} else if strings.HasPrefix(details[i], "Verfügbarab") {
			replacer = "Verfügbarab"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			listing.Availability = details[i]
		} else if strings.HasPrefix(details[i], "Online-Besichtigung") {
			replacer = "Online-Besichtigung"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			listing.OnlineTour = details[i]
		} else if strings.HasPrefix(details[i], "Heizkosten") {
			replacer = "Heizkosten"
			details[i] = strings.Replace(details[i], replacer, "", 1)
			listing.HeatingCosts = details[i]
		}

	}

	return listing
}

// StartEbayCrawl starts the crawling process for ebay-kleinanzeigen
func StartEbayCrawl(url string) {
	// Declaring alternating user agents
	round_robin_ua, _ := NewUserAgent(
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

	if err := c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4}); err != nil {
		fmt.Println(err)
	}
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
		if err := detailCollector.Visit("https://ebay-kleinanzeigen.de" + e.ChildAttr("a.ellipsis", "href")); err != nil {
			fmt.Println(err)
		}
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
		listing = GetDetailsEbay(details, listing)

		// Scraping images
		e.ForEach("img[id=viewad-image]", func(_ int, image *colly.HTMLElement) {
			images = append(images, image.Attr("src"))
		})
		listing.Images = images
		listings = append(listings, listing)
	})

	// Visiting and waiting
	if err := c.Visit(url); err != nil {
		fmt.Println(err)
	}
	c.Wait()

	// Waiting for the jobs to be complete
	time.Sleep(5 * time.Second)

	// Looping through the listings and printing out the resulting objects
	for i := range listings {
		fmt.Println(listings[i])
	}
}
