package services

import (
	"github.com/arangodb/go-driver"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"github.com/yzaimoglu/flathunter/pkg/utils"
)

// GetListing retrieves a listing from the database.
func GetListing(listingId string) (models.Listing, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var listing models.Listing

	result, err := arango.Database.Query(arango.Ctx,
		"FOR listing IN listings FILTER listing._key == @id RETURN listing",
		map[string]interface{}{"id": listingId})
	if err != nil {
		slog.Error(err)
		return models.Listing{}, config.ErrListingNotFound
	}
	defer result.Close()

	_, err = result.ReadDocument(arango.Ctx, &listing)
	if err != nil {
		slog.Errorf("Failed to read document: %v", err)
		return models.Listing{}, config.ErrListingNotFound
	}

	slog.Infof("Retrieved listing with key %s from the database.", listingId)
	return listing, nil
}

// GetListings retrieves the listings from the database.
func GetListings(page int) ([]models.Listing, error) {
	amount := 25
	arango := config.NewArangoClient()
	defer arango.Close()

	var listings []models.Listing

	result, err := arango.Database.Query(arango.Ctx,
		"FOR listing IN listings SORT listing.created_at DESC LIMIT @offset, @limit RETURN listing",
		map[string]interface{}{"offset": (page - 1) * amount, "limit": amount})
	if err != nil {
		slog.Error(err)
		return []models.Listing{}, config.ErrListingNotFound
	}

	defer result.Close()

	for {
		var listing models.Listing
		_, err := result.ReadDocument(arango.Ctx, &listing)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			slog.Errorf("Failed to read document: %v", err)
			return []models.Listing{}, config.ErrListingNotFound
		}
		listings = append(listings, listing)
	}

	if len(listings) == 0 {
		return []models.Listing{}, config.ErrListingNotFound
	}

	slog.Infof("Retrieved %d listings from the database.", len(listings))
	return listings, nil
}

// InsertListing inserts a listing into the database.
func InsertListing(createListing models.Listing, url models.URL) (string, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, "listings")
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return "", err
	}

	meta, err := collection.CreateDocument(arango.Ctx, createListing)
	if err != nil {
		slog.Errorf("Failed to create document: %v", err)
		return "", err
	}

	slog.Infof("Inserted listing with key %s into the database.", meta.Key)
	return meta.Key, nil
}

// InsertListings inserts multiple listings into the database.
func InsertListings(listings []models.Listing, url models.URL) ([]string, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var insertedListings []string

	collection, err := arango.Database.Collection(arango.Ctx, "listings")
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return []string{}, err
	}

	for _, listing := range listings {
		if !ListingExists(listing) {
			insertedId, err := collection.CreateDocument(arango.Ctx, listing)
			if err != nil {
				slog.Errorf("Failed to create document: %v", err)
				return []string{}, err
			}
			insertedListings = append(insertedListings, insertedId.Key)
		}
	}

	slog.Infof("Inserted %d listings into the database.", len(insertedListings))
	return insertedListings, nil
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
func GetUserListings(userId string, page int) ([]models.UserListing, error) {
	amount := 25
	arango := config.NewArangoClient()
	defer arango.Close()

	var userListings []models.UserListing

	query := "FOR user_listing IN user_listings FILTER user_listing.user == @userId FOR user IN users FILTER user._key == user_listing.user FOR listing IN listings FILTER listing._key == user_listing.listing FOR role IN roles FILTER role._key == user.role LIMIT @offset, @limit RETURN merge(user_listing, {user: merge(user, {role: role})}, {listing: listing})"
	bindVars := map[string]interface{}{
		"userId": userId,
		"limit":  amount,
		"offset": (page - 1) * amount,
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

	if len(userListings) == 0 {
		return []models.UserListing{}, config.ErrListingNotFound
	}

	utils.UserListingsToSafe(&userListings)
	return userListings, nil
}

// GetUserListings gets a user listing from the database.
func GetUserListing(userId string, listingId string) (models.UserListing, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var userListing models.UserListing

	query := "FOR user_listing IN user_listings FILTER user_listing.user == @userId && user_listing._key == @listingId FOR user IN users FILTER user._key == user_listing.user FOR listing IN listings FILTER listing._key == user_listing.listing FOR role IN roles FILTER role._key == user.role RETURN merge(user_listing, {user: merge(user, {role: role})}, {listing: listing})"
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

	utils.UserListingToSafe(&userListing)
	return userListing, nil
}

// InsertUserListing inserts a user listing into the database.
func InsertUserListing(createListing models.CreateUserListing) (string, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUserListingsCollection)
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return "", err
	}

	meta, err := collection.CreateDocument(arango.Ctx, createListing)
	if err != nil {
		slog.Errorf("Failed to create document: %v", err)
		return "", err
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
		_, err = collection.CreateDocument(arango.Ctx, listing)
		if err != nil {
			slog.Errorf("Failed to create document: %v", err)
			return err
		}
		insertedNumber++
	}

	slog.Infof("Inserted %d listings into the database.", insertedNumber)
	return nil
}

// DeleteUserListing deletes a user listing from the database.
func DeleteUserListing(listingId string) error {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUserListingsCollection)
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return err
	}

	_, err = collection.RemoveDocument(arango.Ctx, listingId)
	if err != nil {
		slog.Errorf("Failed to remove document: %v", err)
		return err
	}

	slog.Infof("Removed user listing with key %s from the database.", listingId)
	return nil
}
