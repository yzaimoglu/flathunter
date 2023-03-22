package crawler

import "github.com/yzaimoglu/flathunter/pkg/models"

const (
	EbayKleinazeigen = "ebay_kleinanzeigen"
	WgGesucht        = "wg_gesucht"
)

// Crawler is the base struct for all crawlers
type Crawler struct {
	URL       models.URL
	UserAgent *models.UserAgent
	Proxy     *models.Proxy
}

// InitCrawler is the init function for all crawlers
func InitCrawler(url models.URL, ua *models.UserAgent, proxy *models.Proxy) Crawler {
	return Crawler{
		UserAgent: ua,
		Proxy:     proxy,
		URL:       url,
	}
}

// Crawl is the main function for all crawlers
func (crawler Crawler) Crawl() {
	switch crawler.URL.Platform.Name {
	case EbayKleinazeigen:
		StartEbayCrawl(crawler.URL.URL, crawler.UserAgent, crawler.Proxy)
		return
	case WgGesucht:
		StartWgGesuchtCrawl(crawler.URL.URL, crawler.UserAgent, crawler.Proxy)
		return
	}
}
