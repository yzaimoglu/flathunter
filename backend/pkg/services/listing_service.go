package services

import (
	"github.com/arangodb/go-driver"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

// InsertListing inserts a listing into the database.
func InsertListing(createListing models.CreateListing) (interface{}, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, "listings")
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return nil, err
	}

	meta, err := collection.CreateDocument(arango.Ctx, createListing)
	if err != nil {
		slog.Errorf("Failed to create document: %v", err)
		return nil, err
	}

	slog.Infof("Inserted url with key %s into the database.", meta.Key)
	return meta.Key, nil
}

// InsertListings inserts multiple listings into the database.
func InsertListings(listings []models.Listing) error {
	insertedNumber := 0
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, "listings")
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return err
	}

	for _, listing := range listings {
		if !ListingExists(listing) {
			_, err = collection.CreateDocument(arango.Ctx, listing)
			if err != nil {
				slog.Errorf("Failed to create document: %v", err)
				return err
			}
			insertedNumber++
		}
	}

	slog.Infof("Inserted %d listings into the database.", insertedNumber)
	return nil
}

// ListingExists checks if a listing already exists in the database.
func ListingExists(listing models.Listing) bool {
	arango := config.NewArangoClient()
	defer arango.Close()

	var result models.Listing

	query := "FOR listing IN listings FILTER listing.url == @url RETURN listing"
	bindVars := map[string]interface{}{
		"url": listing.URL,
	}

	cursor, err := arango.Database.Query(arango.Ctx, query, bindVars)
	if err != nil {
		slog.Errorf("Failed to query database: %v", err)
		return false
	}
	defer cursor.Close()

	_, err = cursor.ReadDocument(arango.Ctx, &result)
	if driver.IsNoMoreDocuments(err) {
		return false
	} else if err != nil {
		slog.Errorf("Failed to read document: %v", err)
		return false
	}

	return true
}

// GetUserListings gets a user listing from the database.
func GetUserListings(userId string) ([]models.UserListing, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var userListings []models.UserListing

	query := "FOR user_listing IN user_listings FILTER user_listing.user == @userId RETURN listing"
	bindVars := map[string]interface{}{
		"userId": userId,
	}

	result, err := arango.Database.Query(arango.Ctx, query, bindVars)
	if err != nil {
		slog.Errorf("Failed to query database: %v", err)
		return nil, err
	}
	defer result.Close()

	for {
		var listing models.UserListing
		_, err := result.ReadDocument(arango.Ctx, &listing)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			slog.Errorf("Failed to read document: %v", err)
			return nil, err
		}
		userListings = append(userListings, listing)
	}

	return userListings, nil
}

// GetUserListings gets a user listing from the database.
func GetUserListing(userId string, listingId string) (models.UserListing, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var userListing models.UserListing

	query := "FOR user_listing IN user_listings FILTER user_listing.user == @userId && user_listing._key == @listingId RETURN listing"
	bindVars := map[string]interface{}{
		"userId":    userId,
		"listingId": listingId,
	}

	result, err := arango.Database.Query(arango.Ctx, query, bindVars)
	if err != nil {
		slog.Errorf("Failed to query database: %v", err)
		return models.UserListing{}, err
	}
	defer result.Close()

	_, err = result.ReadDocument(arango.Ctx, &userListing)
	if driver.IsNoMoreDocuments(err) {
		return models.UserListing{}, config.ErrURLNotFound
	} else if err != nil {
		slog.Errorf("Failed to read document: %v", err)
		return models.UserListing{}, config.ErrURLNotFound
	}

	return userListing, nil
}

// InsertUserListing inserts a user listing into the database.
func InsertUserListing(createListing models.CreateUserListing) (interface{}, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUserListingsCollection)
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return nil, err
	}

	meta, err := collection.CreateDocument(arango.Ctx, createListing)
	if err != nil {
		slog.Errorf("Failed to create document: %v", err)
		return nil, err
	}

	slog.Infof("Inserted user listing with key %s into the database.", meta.Key)
	return meta.Key, nil
}

// InsertListings inserts multiple listings into the database.
func InsertUserListings(listings []models.UserListing) error {
	insertedNumber := 0
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUserListingsCollection)
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return err
	}

	for _, listing := range listings {
		if !UserListingExists(listing.User.ID, listing.Listing.URL) {
			_, err = collection.CreateDocument(arango.Ctx, listing)
			if err != nil {
				slog.Errorf("Failed to create document: %v", err)
				return err
			}
			insertedNumber++
		}
	}

	slog.Infof("Inserted %d listings into the database.", insertedNumber)
	return nil
}

// UserListingExists checks if a user listing already exists in the database.
func UserListingExists(userId string, listingUrl string) bool {
	return false
}
