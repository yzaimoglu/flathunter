package main

import (
	"time"

	"github.com/alitto/pond"
	"github.com/go-co-op/gocron"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/crawler"
)

func main() {
	round_robin_ua, _ := crawler.GetUserAgentRA()
	round_robin_proxy, _ := crawler.GetProxyRA()
	workerPool := pond.New(150, 1000)

	// Create the crawler application
	crawlerApp := crawler.Crawler{
		WorkerPool:  workerPool,
		ProxyRR:     round_robin_proxy,
		UserAgentRR: round_robin_ua,
	}

	// Load configuration
	config.Load()
	config.SetupLogger()
	config.SetupArango()

	// Setup cron jobs
	cronScheduler := gocron.NewScheduler(time.Now().Location())
	cronScheduler.Every(5).Minute().Do(func() {
		crawlerApp.RunMinuteCron()
	})

	cronScheduler.StartAsync()
	config.SysCallSetup()
	workerPool.StopAndWait()
}
