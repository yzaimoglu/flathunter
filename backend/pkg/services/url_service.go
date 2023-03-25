package services

import (
	"time"

	"github.com/arangodb/go-driver"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

//
// URLS SECTION
//

// GetURLs retrieves all urls to be crawled from the database.
func GetURLs() ([]models.URL, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var urls []models.URL

	result, err := arango.Database.Query(arango.Ctx,
		"FOR url IN urls SORT url.last_crawled ASC LIMIT 50 FILTER url.last_crawled+@crawl_time <= @now FOR platform in platforms FILTER url.platform == platform.name SORT url.last_crawled ASC RETURN merge(url, {platform: platform})",
		map[string]interface{}{"now": (time.Now().Unix()), "crawl_time": 300})
	if err != nil {
		slog.Error(err)
		return []models.URL{}, config.ErrURLsNotFound
	}
	defer result.Close()

	for {
		var url models.URL
		_, err := result.ReadDocument(arango.Ctx, &url)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			slog.Errorf("Failed to read document: %v", err)
			return []models.URL{}, config.ErrURLsNotFound
		}
		urls = append(urls, url)
	}

	if len(urls) == 0 {
		return []models.URL{}, config.ErrURLsNotFound
	}

	slog.Infof("Retrieved %d urls from the database.", len(urls))
	return urls, nil
}

// GetURL retrieves a single url.
func GetURL(id string) (models.URL, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var url models.URL

	result, err := arango.Database.Query(arango.Ctx,
		"FOR url IN urls FILTER url._key == @id FOR platform in platforms FILTER url.platform == platform.name RETURN merge(url, {platform: platform})",
		map[string]interface{}{"id": id})
	if err != nil {
		slog.Error(err)
		return models.URL{}, config.ErrURLNotFound
	}
	defer result.Close()

	_, err = result.ReadDocument(arango.Ctx, &url)
	if driver.IsNoMoreDocuments(err) {
		return models.URL{}, config.ErrURLNotFound
	} else if err != nil {
		slog.Errorf("Failed to read document: %v", err)
		return models.URL{}, config.ErrURLNotFound
	}

	slog.Infof("Retrieved the url with id %s from the database.", id)
	return url, nil
}

// InsertURL inserts a new url into the database.
func InsertURL(createURL models.CreateURLRequest) (interface{}, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, "urls")
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return nil, err
	}

	meta, err := collection.CreateDocument(arango.Ctx, createURL)
	if err != nil {
		slog.Errorf("Failed to create document: %v", err)
		return nil, err
	}

	slog.Infof("Inserted url with key %s into the database.", meta.Key)
	return meta.Key, nil
}

// UpdateURL updates a single field of a url.
func UpdateURL(field string, value interface{}, key string) error {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, "urls")
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return err
	}

	_, err = collection.UpdateDocument(arango.Ctx, key, map[string]interface{}{field: value})
	if err != nil {
		slog.Errorf("Failed to update document: %v", err)
		return err
	}

	slog.Infof("Updated %s of url with key %s in the database.", field, key)
	return nil
}

//
// USER URLS SECTION
//

// TODO: GetUserURL retrieves a single url.
func GetUserURL(userId string, urlId string) (models.URL, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var url models.URL

	result, err := arango.Database.Query(arango.Ctx,
		"FOR url IN user_urls FILTER url._key == @urlId FOR platform in platforms FILTER url.platform == platform.name RETURN merge(url, {platform: platform})",
		map[string]interface{}{"urlId": urlId})
	if err != nil {
		slog.Error(err)
		return models.URL{}, config.ErrURLNotFound
	}
	defer result.Close()

	_, err = result.ReadDocument(arango.Ctx, &url)
	if driver.IsNoMoreDocuments(err) {
		return models.URL{}, config.ErrURLNotFound
	} else if err != nil {
		slog.Errorf("Failed to read document: %v", err)
		return models.URL{}, config.ErrURLNotFound
	}

	slog.Infof("Retrieved the url with id %s of user %s from the database.", urlId, userId)
	return url, nil
}

// TODO: InsertUserURL inserts a user url into the database.
func InsertUserURL(createURL models.CreateUserURL) (interface{}, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUserURLsCollection)
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return nil, err
	}

	meta, err := collection.CreateDocument(arango.Ctx, createURL)
	if err != nil {
		slog.Errorf("Failed to create document: %v", err)
		return nil, err
	}

	slog.Infof("Inserted user url with key %s into the database.", meta.Key)
	return meta.Key, nil
}

// TODO: UpdateURL updates a single field of a url.
func UpdateUserURL(field string, value interface{}, key string) error {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUserURLsCollection)
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return err
	}

	_, err = collection.UpdateDocument(arango.Ctx, key, map[string]interface{}{field: value})
	if err != nil {
		slog.Errorf("Failed to update document: %v", err)
		return err
	}

	slog.Infof("Updated %s of user url with key %s in the database.", field, key)
	return nil
}

// TODO: SetLastCrawledURL updates the last_crawled field of a url.
func SetLastCrawledURL(key string) error {
	return UpdateURL("last_crawled", time.Now().Unix(), key)
}

// TODO: SetLastCrawledURL updates the last_crawled field of a url.
func SetLastCrawledUserURL(key string) error {
	return UpdateUserURL("last_crawled", time.Now().Unix(), key)
}
