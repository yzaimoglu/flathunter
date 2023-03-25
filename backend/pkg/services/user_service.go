package services

import (
	"github.com/arangodb/go-driver"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"github.com/yzaimoglu/flathunter/pkg/utils"
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
func GetUserByID(id string) (models.User, error) {
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

// GetUser retrieves a user from the database.
func GetUserByEmail(email string) (models.User, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var user models.User

	result, err := arango.Database.Query(arango.Ctx,
		"FOR u IN users FILTER u.email == @email FOR r in roles FILTER u.role == r._key RETURN merge(u, {role: r})",
		map[string]interface{}{"email": email})
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

// GetUserSessions retrieves the user sessions from the database.
func GetUserSessions(userID string) ([]models.SessionToken, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var sessions []models.SessionToken

	result, err := arango.Database.Query(arango.Ctx,
		"FOR s IN user_sessions FILTER s.user == @userID RETURN s",
		map[string]interface{}{"userID": userID})
	if err != nil {
		slog.Error(err)
		return []models.SessionToken{}, config.ErrUserNotFound
	}
	defer result.Close()

	for {
		var session models.SessionToken
		_, err := result.ReadDocument(arango.Ctx, &session)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			slog.Errorf("Failed to read document: %v", err)
			return []models.SessionToken{}, config.ErrUserNotFound
		}
		sessions = append(sessions, session)
	}

	slog.Infof("Retrieved %d sessions from the database for user %s.", len(sessions), userID)
	return sessions, nil
}

// InsertUser inserts a new user into the database.
func InsertUser(createUser models.CreateUserRequest) (interface{}, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUsersCollection)
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return nil, err
	}

	// Create a new user.
	user := models.CreateUser{
		Email:          createUser.Email,
		HashedPassword: utils.HashPassword(createUser.Password),
		Role:           "user",
	}

	meta, err := collection.CreateDocument(arango.Ctx, user)
	if err != nil {
		slog.Errorf("Failed to create document: %v", err)
		return nil, err
	}

	slog.Infof("Inserted user with key %s into the database.", meta.Key)
	return meta.Key, nil
}
