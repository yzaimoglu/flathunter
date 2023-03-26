package main

import (
	"github.com/yzaimoglu/flathunter/pkg/models"
	"github.com/yzaimoglu/flathunter/pkg/notifier"
)

func main() {
	// round_robin_ua, _ := crawler.GetUserAgentRA()
	// round_robin_proxy, _ := crawler.GetProxyRA()
	// workerPool := pond.New(150, 1000)

	// // Create the crawler application
	// crawlerApp := crawler.Crawler{
	// 	WorkerPool:  workerPool,
	// 	ProxyRR:     round_robin_proxy,
	// 	UserAgentRR: round_robin_ua,
	// }

	// // Load configuration
	// config.Load()
	// config.SetupLogger()
	// config.SetupArango()

	// // Setup cron jobs
	// cronScheduler := gocron.NewScheduler(time.Now().Location())
	// cronScheduler.Every(3).Minute().Do(func() {
	// 	crawlerApp.RunThreeMinuteCron()
	// })
	// cronScheduler.Every(5).Minute().Do(func() {
	// 	crawlerApp.RunFiveMinuteCron()
	// })

	// cronScheduler.StartAsync()
	// config.SysCallSetup()
	// workerPool.StopAndWait()

	notifier.NotifyDiscord(models.UserListing{
		User: models.User{
			Email:          "email",
			HashedPassword: "hashed_password",
		},
		Listing: models.Listing{
			URL:         "https://www.ebay-kleinanzeigen.de/s-anzeige/3-zimmerwohnung-mit-blick-ins-wandsbeker-gehoelz/2372702682-203-23510",
			Images:      []string{"image1", "image2"},
			Price:       "100€",
			Date:        "2021-01-01",
			Description: "description",
			Size:        "100m²",
			Rooms:       3,
			Bathrooms:   1,
			Floor:       "1",
			Type:        "type",
			ExtraCosts:  "extra_costs",
			FullRent:    "full_rent",
		},
	}, models.Notifier{
		Type: "discord",
		Options: map[string]interface{}{
			"webhook": "https://discord.com/api/webhooks/1089690725488267395/pvnYvdZWU0FJUL9Ye9l0K7mekA4Bw8-Yi2yDh-0WFEuJlJ_kiybzKORhNIMWEe1yjyFk",
		},
	})
}
