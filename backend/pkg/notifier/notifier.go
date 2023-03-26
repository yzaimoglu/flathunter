package notifier

import "github.com/yzaimoglu/flathunter/pkg/models"

// Notify sends a notification to the user
func Notify(listing models.UserListing, notifier models.Notifier) {
	switch notifier.Type {
	case "email":
		NotifyEmail(listing, notifier)
	case "discord":
		NotifyDiscord(listing, notifier)
	}
}
