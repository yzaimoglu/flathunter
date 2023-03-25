package crawler

import (
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"github.com/yzaimoglu/flathunter/pkg/services"
)

// RunMinuteCron runs every minute.
// It gets the last 50 urls from the database and crawls them.
func (crawler Crawler) RunMinuteCron() {
	urls, err := services.GetURLs()
	if err != nil {
		slog.Errorf("Error while getting urls: %s", err.Error())
	}

	for _, url := range urls {
		temp := url
		crawler.WorkerPool.Submit(func() {
			crawlerClient := crawler.InitCrawler(temp, crawler.UserAgentRR.Next(), crawler.ProxyRR.Next())
			listings, err := crawlerClient.Crawl()
			if err != nil {
				slog.Errorf("Error while crawling: %s", err.Error())
				return
			}
			services.SetLastCrawledURL(temp.Key)
			crawler.WorkerPool.Submit(func() {
				listingIds, err := services.InsertListings(listings, temp)
				if err != nil {
					slog.Errorf("Error while inserting listings: %s", err.Error())
					return
				}
				users, err := services.GetUsersByURL(temp.URL)
				if err != nil {
					slog.Errorf("Error while getting users: %s", err.Error())
					return
				}
				for _, user := range users {
					for _, listingId := range listingIds {
						crawler.WorkerPool.Submit(func() {
							_, err := services.InsertUserListing(models.CreateUserListing{
								User:     user.Key,
								Listing:  listingId,
								Notified: false,
							})
							if err != nil {
								slog.Errorf("Error while inserting user listing: %s", err.Error())
								return
							}
						})
					}
				}
			})
		})
	}
}
