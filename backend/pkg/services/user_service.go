package services

import (
	"github.com/arangodb/go-driver"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

// GetUsers retrieves the users from the database.
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

	slog.Infof("Retrieved %d users from the database.", len(users))
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

	slog.Infof("Retrieved user with key %s from the database.", user.Key)
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

	slog.Infof("Inserted user with key %s into the database.", meta.Key)
	return meta.Key, nil
}
