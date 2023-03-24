package crawler

import (
	"github.com/gocolly/colly/v2"
)

// setHeaders sets the headers for the request.
func setHeaders(r *colly.Request, host string, ua string) {
	r.Headers.Set("User-Agent", ua)
	r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
	r.Headers.Set("Accept-Language", "en-GB; en-US,en;q=0.9,de;q=0.8")
	r.Headers.Set("Cache-Control", "max-age=0")
	r.Headers.Set("Connection", "keep-alive")
	r.Headers.Set("Host", host)
	r.Headers.Set("Sec-Fetch-Dest", "document")
	r.Headers.Set("Sec-Fetch-Mode", "navigate")
	r.Headers.Set("Sec-Fetch-Site", "same-origin")
	r.Headers.Set("Sec-Fetch-User", "?1")
	r.Headers.Set("Upgrade-Insecure-Requests", "1")
}
