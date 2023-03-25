package crawler

import (
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

// StartWgGesuchtCrawl is the function to start the crawling process
func StartWgGesuchtCrawl(url string, ua *models.UserAgent, proxy *models.Proxy) ([]models.Listing, error) {
	url = strings.Replace(url, "1.1.0.html", "1.0.0.html", -1)
	var listings []models.Listing = []models.Listing{}

	c := colly.NewCollector(
		colly.UserAgent(ua.UserAgent),
		colly.AllowURLRevisit(),
		//colly.CacheDir("./cache"),
		colly.MaxDepth(2),
		colly.Async(true),
	)

	// Setting the limit for the parallelism
	if err := c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 4}); err != nil {
		slog.Errorf("Could not set the limit for parallelism: %v", err)
	}
	c.SetRequestTimeout(120 * time.Second)

	// Setting proxy
	if config.GetBoolean("ALTERNATING_PROXY") {
		c.SetProxy(ProxyString(proxy))
	}

	// Cloning the colly collector for the detailCollector
	detailCollector := c.Clone()

	// Setting the alternating User Agent and standard Headers
	c.OnRequest(func(r *colly.Request) {
		setHeaders(r, "wg-gesucht.de", ua.UserAgent)
	})
	detailCollector.OnRequest(func(r *colly.Request) {
		setHeaders(r, "wg-gesucht.de", ua.UserAgent)
	})

	// Visiting the listings specific urls to scrape
	c.OnHTML("tr.listenansicht0.offer_list_item", func(e *colly.HTMLElement) {
		link := "https://wg-gesucht.de/" + e.ChildAttr("td.ang_spalte_datum.row_click > a.list.listenansicht0", "href")
		detailCollector.Visit(link)
	})

	// Error while scraping
	c.OnError(func(r *colly.Response, e error) {
		slog.Errorf("Request URL: %v failed with response: %v", r.Request.URL, r.StatusCode)
	})

	// Scraping details
	detailCollector.OnHTML("div[id=main_content] div.panel.panel-default div.panel-body", func(e *colly.HTMLElement) {
		// Creating a new listing
		var listing models.Listing = models.Listing{
			CreatedAt: time.Now().Unix(),
		}

		// Saving the url
		listing.URL = e.Request.URL.String()

		// Scraping description
		listing.Description = strings.TrimSpace(e.ChildText("div[id=ad_description_text] p.freitext"))

		// Scraping images
		var images []string
		e.ForEach("div[id=bildContainer] img.sp-image", func(_ int, image *colly.HTMLElement) {
			images = append(images, image.Attr("data-large"))
		})
		listing.Images = images

		// Scraping row by row to get the details
		e.ForEach("div.row", func(i int, e *colly.HTMLElement) {
			switch i {
			case 2:
				e.ForEach("h2", func(i int, e *colly.HTMLElement) {
					switch i {
					case 0:
						listing.Size = strings.TrimSpace(e.Text)
					case 1:
						listing.FullRent = strings.TrimSpace(e.Text)
					}
				})
			case 4:
				e.ForEach("div.col-sm-3 p b", func(i int, e *colly.HTMLElement) {
					switch i {
					// Available from
					case 0:
						listing.Availability = strings.TrimSpace(e.Text)
					// Available till
					case 1:
					// Online since
					case 2:
					}
				})

				// Price, HeatingCosts, ExtraCosts & Deposit
				e.ForEach("div", func(i int, e *colly.HTMLElement) {
					if i == 0 {
						e.ForEach("table", func(i int, e *colly.HTMLElement) {
							e.ForEach("td", func(i int, e *colly.HTMLElement) {
								switch i {
								case 1:
									listing.Price = strings.TrimSpace(e.Text)
								case 3:
									listing.HeatingCosts = strings.TrimSpace(e.Text)
								case 5:
									listing.ExtraCosts = strings.TrimSpace(e.Text)
								case 7:
									listing.Deposit = strings.TrimSpace(e.Text)
								}
							})
						})
					}
				})

			}
		})
		listings = append(listings, listing)
	})

	// Visiting and waiting
	if err := c.Visit(url); err != nil {
		slog.Errorf("Error while visiting the url: %s", err)
		return []models.Listing{}, err
	}
	c.Wait()
	time.Sleep(10 * time.Second)

	slog.Infof("Successfully scraped %d listings from %s", len(listings), url)
	return listings, nil
}
