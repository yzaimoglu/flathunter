package crawler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

// GetDetailsWGGesucht is the function to get the details of a listing
func GetDetailsWGGesucht(details []string, listing models.Listing) (resultingListing models.Listing) {
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

// StartWgGesuchtCrawl is the function to start the crawling process
func StartWgGesuchtCrawl(url string, ua *models.UserAgent, proxy *models.Proxy) ([]models.Listing, error) {
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
		fmt.Println(err)
	}
	c.SetRequestTimeout(120 * time.Second)

	// Setting proxy
	// c.SetProxy(ProxyString(proxy))

	// Cloning the colly collector for the detailCollector
	detailCollector := c.Clone()

	// Setting the alternating User Agent and standard Headers
	c.OnRequest(func(r *colly.Request) {
		setHeaders(r, "wg-gesucht.de", ua.UserAgent)
		r.Headers.Set("Cookie", "PHPSESSID=99jkru1bsibstdb47h2i8mfeq9; X-Client-Id=wg_desktop_website; sync_favourites=false; get_favourites=false; __cmpcpcx15144=__1__; __cmpcpc=__1__; __cmpconsentx15144=CPpIINgPpIINgAfR4BENC9CsAP_AAH_AAAYgG7pV9W5WTWFBOHp7arsEKYUX13TNQ2AiCgCAE6AAiHKAYIQGkmAYJASAIAACIBAgIBYBIQFAAEFEAAAAIIARAAFIAAAAIAAIIAIECAEQUkAAAAAIAAAAAAAAAAAEABAAgADAABIAAEAAAIAAAAAAAAgbulX1blZNYUE4entquwQphRfXdM1DYCIKAIAToACIcoBghAaSYBgkBIAgAAIgECAgFgEhAUAAQUQAAAAggBEAAUgAAAAgAAggAgQIARBSQAAAAAgAAAAAAAAAAAQAEACAAMAAEgAAQAAAgAAAAAAACAAA; __cmpcvcx15144=__s1227_s87_s343_s94_s40_s1052_s64_s335_s914_s762_s640_s102_s1469_s405_s1932_s65_s23_s209_s116_s25_s56_s123_s127_s570_s128_s7_s573_s482_s312_s1_s26_s2612_s135_s1409_s905_s10_s139_s161_s1442_s2_s974_s1049_s11_s322_s2386_s885_s879_s36_s1358_s267_s883_s1097_s2589_s76_c4566_s1341_s268_s460_s271_c13455_s292_s358_s190_s19_s653_s800_s12_s196_s1216_s52_s199_s34_s525_s32_s882_s739_s60_s21_c5169_s35_s30_s217_s574_s356_U__; __cmpcvc=__s1227_s87_s343_s94_s40_s1052_s64_s335_s914_s762_s640_s102_s1469_s405_s1932_s65_s23_s209_s116_s25_s56_s123_s127_s570_s128_s7_s573_s482_s312_s1_s26_s2612_s135_s1409_s905_s10_s139_s161_s1442_s2_s974_s1049_s11_s322_s2386_s885_s879_s36_s1358_s267_s883_s1097_s2589_s76_c4566_s1341_s268_s460_s271_c13455_s292_s358_s190_s19_s653_s800_s12_s196_s1216_s52_s199_s34_s525_s32_s882_s739_s60_s21_c5169_s35_s30_s217_s574_s356_U__; __cmpiab=_58_272_40_231_147_44_50_790_39_14_93_511_612_264_565_6_410_211_195_259_793_23_728_394_742_63_771_273_156_12_87_128_185_30_94_620_315_243_285_77_138_591_85_91_541_440_209_397_122_144_126_434_584_402_8_213_141_183_24_312_1_120_78_755_98_61_206_131_365_606_253_10_278_428_436_129_252_294_62_325_148_97_109_95_508_486_52_614_142_79_152_358_151_20_130_812_373_304_241_617_602_69_227_349_385_772_559_164_412_384_140_490_177_236_887_76_81_835_11_71_4_16_506_84_33_111_73_68_82_161_45_115_134_295_104_13_655_165_238_137_136_114_275_42_89_475_132_345_577_382_21_28_36_162_237_284_18_281_32_25_70_173_154_210_301_469_; __cmpiabc=__1_2_3_4_5_6_7_8_9_10_r1_r2_; __cmpiabli=_58_272_40_231_147_44_50_790_39_14_93_511_612_264_565_6_410_211_195_259_793_23_728_394_742_63_771_273_156_12_87_128_185_30_94_620_315_243_285_77_138_591_85_91_541_440_209_397_122_144_126_434_584_402_8_213_141_183_24_312_1_120_78_755_98_61_206_131_365_606_253_10_278_428_436_129_252_294_62_325_148_97_109_95_508_486_52_614_142_79_152_358_151_20_130_812_373_304_241_617_602_69_227_349_385_772_559_164_412_384_140_490_177_236_887_76_81_835_11_71_4_16_506_84_33_111_73_68_82_161_45_115_134_295_104_13_655_165_238_137_136_114_275_42_89_475_132_345_577_382_21_28_36_162_237_284_18_281_32_25_70_173_154_210_301_469_772_; __cmpiabcli=__2_3_4_5_6_7_8_9_10_; last_city=35; last_cat=1; last_type=0")
	})
	detailCollector.OnRequest(func(r *colly.Request) {
		setHeaders(r, "wg-gesucht.de", ua.UserAgent)
		r.Headers.Set("Cookie", "PHPSESSID=99jkru1bsibstdb47h2i8mfeq9; X-Client-Id=wg_desktop_website; sync_favourites=false; get_favourites=false; __cmpcpcx15144=__1__; __cmpcpc=__1__; __cmpconsentx15144=CPpIINgPpIINgAfR4BENC9CsAP_AAH_AAAYgG7pV9W5WTWFBOHp7arsEKYUX13TNQ2AiCgCAE6AAiHKAYIQGkmAYJASAIAACIBAgIBYBIQFAAEFEAAAAIIARAAFIAAAAIAAIIAIECAEQUkAAAAAIAAAAAAAAAAAEABAAgADAABIAAEAAAIAAAAAAAAgbulX1blZNYUE4entquwQphRfXdM1DYCIKAIAToACIcoBghAaSYBgkBIAgAAIgECAgFgEhAUAAQUQAAAAggBEAAUgAAAAgAAggAgQIARBSQAAAAAgAAAAAAAAAAAQAEACAAMAAEgAAQAAAgAAAAAAACAAA; __cmpcvcx15144=__s1227_s87_s343_s94_s40_s1052_s64_s335_s914_s762_s640_s102_s1469_s405_s1932_s65_s23_s209_s116_s25_s56_s123_s127_s570_s128_s7_s573_s482_s312_s1_s26_s2612_s135_s1409_s905_s10_s139_s161_s1442_s2_s974_s1049_s11_s322_s2386_s885_s879_s36_s1358_s267_s883_s1097_s2589_s76_c4566_s1341_s268_s460_s271_c13455_s292_s358_s190_s19_s653_s800_s12_s196_s1216_s52_s199_s34_s525_s32_s882_s739_s60_s21_c5169_s35_s30_s217_s574_s356_U__; __cmpcvc=__s1227_s87_s343_s94_s40_s1052_s64_s335_s914_s762_s640_s102_s1469_s405_s1932_s65_s23_s209_s116_s25_s56_s123_s127_s570_s128_s7_s573_s482_s312_s1_s26_s2612_s135_s1409_s905_s10_s139_s161_s1442_s2_s974_s1049_s11_s322_s2386_s885_s879_s36_s1358_s267_s883_s1097_s2589_s76_c4566_s1341_s268_s460_s271_c13455_s292_s358_s190_s19_s653_s800_s12_s196_s1216_s52_s199_s34_s525_s32_s882_s739_s60_s21_c5169_s35_s30_s217_s574_s356_U__; __cmpiab=_58_272_40_231_147_44_50_790_39_14_93_511_612_264_565_6_410_211_195_259_793_23_728_394_742_63_771_273_156_12_87_128_185_30_94_620_315_243_285_77_138_591_85_91_541_440_209_397_122_144_126_434_584_402_8_213_141_183_24_312_1_120_78_755_98_61_206_131_365_606_253_10_278_428_436_129_252_294_62_325_148_97_109_95_508_486_52_614_142_79_152_358_151_20_130_812_373_304_241_617_602_69_227_349_385_772_559_164_412_384_140_490_177_236_887_76_81_835_11_71_4_16_506_84_33_111_73_68_82_161_45_115_134_295_104_13_655_165_238_137_136_114_275_42_89_475_132_345_577_382_21_28_36_162_237_284_18_281_32_25_70_173_154_210_301_469_; __cmpiabc=__1_2_3_4_5_6_7_8_9_10_r1_r2_; __cmpiabli=_58_272_40_231_147_44_50_790_39_14_93_511_612_264_565_6_410_211_195_259_793_23_728_394_742_63_771_273_156_12_87_128_185_30_94_620_315_243_285_77_138_591_85_91_541_440_209_397_122_144_126_434_584_402_8_213_141_183_24_312_1_120_78_755_98_61_206_131_365_606_253_10_278_428_436_129_252_294_62_325_148_97_109_95_508_486_52_614_142_79_152_358_151_20_130_812_373_304_241_617_602_69_227_349_385_772_559_164_412_384_140_490_177_236_887_76_81_835_11_71_4_16_506_84_33_111_73_68_82_161_45_115_134_295_104_13_655_165_238_137_136_114_275_42_89_475_132_345_577_382_21_28_36_162_237_284_18_281_32_25_70_173_154_210_301_469_772_; __cmpiabcli=__2_3_4_5_6_7_8_9_10_; last_city=35; last_cat=1; last_type=0")
	})

	// Visiting the listings specific urls to scrape
	c.OnHTML("div[id=main_column]", func(e *colly.HTMLElement) {
		e.ForEach("div.wgg_card.offer_list_item", func(i int, e *colly.HTMLElement) {
			e.ForEach("a", func(i int, e *colly.HTMLElement) {
				link := e.Attr("href")
				if i == 0 {
					detailCollector.Visit("https://wg-gesucht.de" + link)
					slog.Infof("Visiting link: %s", link)
				}
			})
		})
	})

	// c.OnHTML("div", func(e *colly.HTMLElement) {
	// 	fmt.Println(e)
	// })

	// Error while scraping
	c.OnError(func(r *colly.Response, e error) {
		slog.Errorf("Request URL: %v failed with response: %v", r.Request.URL, r.StatusCode)
	})

	// Scraping details
	detailCollector.OnHTML("div.panel-body", func(e *colly.HTMLElement) {
		// Creating a new listing
		var listing models.Listing = models.Listing{
			CreatedAt: time.Now().Unix(),
		}

		// Saving the url
		listing.URL = e.Request.URL.String()

		// Scraping images
		var images []string
		e.ForEach("div[id=bildContainer] img.sp-image", func(_ int, image *colly.HTMLElement) {
			images = append(images, image.Attr("data-large"))
			fmt.Println(images)
		})
		listing.Images = images

		// Scraping row by row to get the details
		e.ForEach("div.row", func(i int, e *colly.HTMLElement) {
			switch i {
			case 0:
			case 1:
				e.ForEach("h2", func(i int, e *colly.HTMLElement) {
					switch i {
					case 0:
						listing.Size = strings.TrimSpace(e.Text)
					case 1:
						listing.FullRent = strings.TrimSpace(e.Text)
					}
				})
			case 3:
				e.ForEach("div", func(i int, e *colly.HTMLElement) {
					switch i {
					case 0:
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
					case 2:
						e.ForEach("p", func(i int, e *colly.HTMLElement) {
							fmt.Println(e)
							switch i {
							case 0:
								listing.Availability = strings.TrimSpace(e.Text)
								fmt.Println(strings.TrimSpace(e.Text))
							}
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

	time.Sleep(3 * time.Second)
	slog.Infof("Successfully scraped %d listings from %s", len(listings), url)
	return listings, nil
}
