package notifier

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

// NotifyDiscord sends a Discord notification to the user
func NotifyDiscord(listing models.UserListing, notifier models.Notifier) {
	body := "{\"content\":null,\"embeds\":[{\"title\":\"$$title$$\",\"description\":\"**Datum:** $$date$$\\n**Preis:** $$price$$\\n**Zimmer:** $$rooms$$\\n**Weitere Details:** [$$url$$]($$url$$)\\n\\n**Beschreibung:**\\n$$description$$\",\"color\":7547331,\"footer\":{\"text\":\"© Flathunter\"},\"image\":{\"url\":\"$$image_url$$\"}}],\"username\":\"Flathunter\",\"avatar_url\":\"$$avatar_url$$\"}"

	if len(listing.Listing.Images) == 0 || !strings.HasPrefix(listing.Listing.Images[0], "https://") {
		slog.Warnf("No image found for listing %s", listing.Listing.URL)
		listing.Listing.Images = []string{"https://img.ebay-kleinanzeigen.de/api/v1/prod-ads/images/f9/f9524e4d-dd8f-431c-848c-f9e09153bf9a?rule=$_59.JPG"}
	}

	// Replace placeholders with actual values
	body = strings.ReplaceAll(body, "$$title$$", "Flathunter - Neue Wohnung gefunden!")
	body = strings.ReplaceAll(body, "$$date$$", listing.Listing.Date)
	body = strings.ReplaceAll(body, "$$price$$", listing.Listing.Price)
	body = strings.ReplaceAll(body, "$$rooms$$", strconv.Itoa(listing.Listing.Rooms))
	body = strings.ReplaceAll(body, "$$url$$", listing.Listing.URL)
	body = strings.ReplaceAll(body, "$$description$$", listing.Listing.Description)
	body = strings.ReplaceAll(body, "$$image_url$$", listing.Listing.Images[0])
	body = strings.ReplaceAll(body, "$$avatar_url$$", "https://raw.githubusercontent.com/yzaimoglu/flathunter/main/backend/assets/logo_symbol.jpg")

	// Send request
	optionsJSON, err := json.Marshal(notifier.Options)
	if err != nil {
		slog.Errorf("Error while marshalling webhook url: %s", err)
		return
	}

	optionsJSONObject := make(map[string]interface{})
	err = json.Unmarshal(optionsJSON, &optionsJSONObject)
	if err != nil {
		slog.Errorf("Error while unmarshalling webhook url: %s", err)
		return
	}
	webhookURL := optionsJSONObject["webhook"]

	_, err = http.Post(webhookURL.(string), "application/json", strings.NewReader(body))
	if err != nil {
		slog.Errorf("Error while sending notification to Discord: %s", err)
	}

	// defer res.Body.Close()
	// responseBody, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(res.Status)
	// fmt.Println(string(responseBody))

	slog.Infof("Sent Discord notification for listing %s and user %s", listing.Listing.URL, listing.User.Key)
}
