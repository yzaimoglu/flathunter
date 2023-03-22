package services

import (
	"github.com/arangodb/go-driver"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

// GetUser retrieves a user from the database.
func GetUsers() ([]models.User, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var users []models.User

	result, err := arango.Database.Query(arango.Ctx,
		"FOR u IN users FOR r in roles FILTER u.role == r._key RETURN merge(u, {role: r})",
		map[string]interface{}{})
	if err != nil {
		slog.Error(err)
		return []models.User{}, config.ErrUserNotFound
	}
	defer result.Close()

	for {
		var user models.User
		_, err := result.ReadDocument(arango.Ctx, &user)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			slog.Errorf("Failed to read document: %v", err)
			return []models.User{}, config.ErrUserNotFound
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUser retrieves a user from the database.
func GetUser(id string) (models.User, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var user models.User

	result, err := arango.Database.Query(arango.Ctx,
		"FOR u IN users FILTER u._key == @id FOR r in roles FILTER u.role == r._key RETURN merge(u, {role: r})",
		map[string]interface{}{"id": id})
	if err != nil {
		slog.Error(err)
		return models.User{}, config.ErrUserNotFound
	}
	defer result.Close()

	_, err = result.ReadDocument(arango.Ctx, &user)
	if driver.IsNoMoreDocuments(err) || err != nil {
		slog.Errorf("Failed to read document: %v", err)
		return models.User{}, config.ErrUserNotFound
	}

	return user, nil
}

// InsertUser inserts a new user into the database.
func InsertUser(createUser models.CreateUser) (interface{}, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, "users")
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return nil, err
	}

	meta, err := collection.CreateDocument(arango.Ctx, createUser)
	if err != nil {
		slog.Errorf("Failed to create document: %v", err)
		return nil, err
	}

	return meta.Key, nil
}
