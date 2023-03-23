package crawler

import (
	"github.com/alitto/pond"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

const (
	EbayKleinazeigen = "ebay_kleinanzeigen"
	WgGesucht        = "wg_gesucht"
)

// Crawler is the base struct for the crawler application
type Crawler struct {
	WorkerPool  *pond.WorkerPool
	ProxyRR     RoundRobinProxy
	UserAgentRR RoundRobinUA
}

// CrawlerClient is the base struct for all crawlers
type CrawlerClient struct {
	URL       models.URL
	UserAgent *models.UserAgent
	Proxy     *models.Proxy
}

// InitCrawler is the init function for all crawlers
func (crawler Crawler) InitCrawler(url models.URL, ua *models.UserAgent, proxy *models.Proxy) CrawlerClient {
	return CrawlerClient{
		UserAgent: ua,
		Proxy:     proxy,
		URL:       url,
	}
}

// Crawl is the main function for all crawlers
func (crawler CrawlerClient) Crawl() []models.Listing {
	var listings []models.Listing
	switch crawler.URL.Platform.Name {
	case EbayKleinazeigen:
		listings = StartEbayCrawl(crawler.URL.URL, crawler.UserAgent, crawler.Proxy)
		return listings
	case WgGesucht:
		listings = StartWgGesuchtCrawl(crawler.URL.URL, crawler.UserAgent, crawler.Proxy)
		return listings
	}

	return []models.Listing{}
}
