package crawler

import (
	"github.com/gookit/slog"
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
			services.SetLastCrawledURL(temp.Key)
			crawlerClient := crawler.InitCrawler(temp, crawler.UserAgentRR.Next(), crawler.ProxyRR.Next())
			listings := crawlerClient.Crawl()
			crawler.WorkerPool.Submit(func() {
				services.InsertListings(listings)
			})
		})
	}
}
