package main

import (
	"fmt"
	"time"

	"github.com/alitto/pond"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/crawler"
	"github.com/yzaimoglu/flathunter/pkg/http/services"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

func main() {
	round_robin_ua, _ := crawler.GetUserAgentRA()
	round_robin_proxy, _ := crawler.GetProxyRA()
	workerPool := pond.New(150, 1000)

	config.Load()
	config.SetupLogger()
	config.SetupArango()

	fmt.Println(time.Now().Unix() + 60*5)

	for i := 0; i < 50; i++ {
		fmt.Println(round_robin_ua.Next())
		fmt.Println(round_robin_proxy.Next())
	}

	for i := 0; i < 500; i++ {
		workerPool.Submit(func() {
			// crawler.StartEbayCrawl("http://localhost:50000/api/v1/users/get", round_robin_ua.Next(), round_robin_proxy.Next())
			services.InsertURL(models.CreateURL{
				URL:         "https://ebay-kleinanzeigen.de/" + fmt.Sprint(time.Now().Unix()+int64(i)),
				Platform:    "ebay_kleinanzeigen",
				LastCrawled: time.Now().Unix(),
				CreatedAt:   time.Now().Unix(),
			})
		})
	}

	workerPool.StopAndWait()
}
