package utils

import (
	"github.com/yzaimoglu/flathunter/pkg/models"
)

// UserToSafe removes the hashed password from the user
func UserToSafe(user *models.User) {
	user.HashedPassword = ""
}

// UsersToSafe removes the hashed password from the user
func UsersToSafe(users *[]models.User) {
	for i := range *users {
		UserToSafe(&(*users)[i])
	}
}

// UserListingToSafe removes the hashed password from the user
func UserListingToSafe(listing *models.UserListing) {
	listing.User.HashedPassword = ""
}

// UserListingsToSafe removes the hashed password from the user
func UserListingsToSafe(listings *[]models.UserListing) {
	for i := range *listings {
		UserListingToSafe(&(*listings)[i])
	}
}

// UserURLToSafe removes the hashed password from the user
func UserURLToSafe(url *models.UserURL) {
	url.User.HashedPassword = ""
}

// UserURLsToSafe removes the hashed password from the user
func UserURLsToSafe(urls *[]models.UserURL) {
	for i := range *urls {
		UserURLToSafe(&(*urls)[i])
	}
}

// SessionToSafe removes the hashed password from the user
func SessionToSafe(session *models.SessionToken) {
	session.User.HashedPassword = ""
}

// SessionsToSafe removes the hashed password from the user
func SessionsToSafe(sessions *[]models.SessionToken) {
	for i := range *sessions {
		SessionToSafe(&(*sessions)[i])
	}
}
