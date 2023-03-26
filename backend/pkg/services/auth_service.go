package services

import (
	"github.com/arangodb/go-driver"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"github.com/yzaimoglu/flathunter/pkg/utils"
)

// LoginUser logs in the user
func LoginUser(loginUser models.LoginUserRequest) (models.SessionToken, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	// Check if the user exists and get the user
	user, err := GetUserByEmail(loginUser.Email)
	if err != nil {
		return models.SessionToken{}, config.ErrInvalidEmail
	}

	// Check if the password is correct
	if !utils.CheckPassword(user.HashedPassword, loginUser.Password) {
		return models.SessionToken{}, config.ErrInvalidPassword
	}

	// Creta a new session token
	sessionToken, err := utils.CreateSession(user)
	if err != nil {
		return models.SessionToken{}, config.ErrSessionCreation
	}

	// Create a new session in the database
	collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUserSessionsCollection)
	if err != nil {
		slog.Errorf("Failed to retrieve collection: %v", err)
		return models.SessionToken{}, err
	}

	_, err = collection.CreateDocument(arango.Ctx, models.CreateSessionToken{
		User:         user.Key,
		SessionToken: sessionToken.SessionToken,
		CreatedAt:    sessionToken.CreatedAt,
		ExpiresAt:    sessionToken.ExpiresAt,
	})
	if err != nil {
		slog.Errorf("Failed to create document: %v", err)
		return models.SessionToken{}, err
	}

	utils.SessionToSafe(&sessionToken)
	slog.Infof("User %s logged in.", user.Email)
	return sessionToken, nil
}

// LogoutUser logs out the user
func LogoutUserAllSessions(logoutUser models.User) error {
	arango := config.NewArangoClient()
	defer arango.Close()

	// Get all sessions of the user
	sessions, err := GetUserSessions(logoutUser.Key)
	if err != nil {
		return err
	}

	// Delete all sessions of the user
	for _, session := range sessions {
		collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUserSessionsCollection)
		if err != nil {
			return err
		}
		_, err = collection.RemoveDocument(arango.Ctx, session.Key)
		if err != nil {
			return err
		}
	}

	slog.Infof("User %s logged out from all %d sessions.", logoutUser.Key, len(sessions))
	return nil
}

// LogoutUser logs out the user
func LogoutUser(sessionToken string) error {
	arango := config.NewArangoClient()
	defer arango.Close()

	// Delete the session
	result, err := arango.Database.Query(arango.Ctx,
		"FOR s IN user_sessions FILTER s.session_token == @token REMOVE s IN user_sessions",
		map[string]interface{}{"token": sessionToken})
	if err != nil {
		slog.Error(err)
		return err
	}
	defer result.Close()

	slog.Infof("Session %s was deleted.", sessionToken)
	return nil
}

// GetSessionByToken retrieves the session from the database.
func GetSessionByToken(sessionToken string) (models.SessionToken, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var sessionObject models.SessionToken

	result, err := arango.Database.Query(arango.Ctx,
		"FOR s IN user_sessions FILTER s.session_token == @token FOR u in users FILTER u._key == s.user RETURN s",
		map[string]interface{}{"token": sessionToken})
	if err != nil {
		slog.Error(err)
		return models.SessionToken{}, config.ErrSessionNotFound
	}
	defer result.Close()

	_, err = result.ReadDocument(arango.Ctx, &sessionObject)
	if driver.IsNoMoreDocuments(err) || err != nil {
		slog.Errorf("Failed to read document: %v", err)
		return models.SessionToken{}, config.ErrSessionNotFound
	}

	utils.SessionToSafe(&sessionObject)
	return sessionObject, nil
}

// GetSessionWithUserByToken retrieves the session with the user from the database.
func GetSessionWithUserByToken(sessionToken string) (models.SessionToken, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var sessionObject models.SessionToken

	result, err := arango.Database.Query(arango.Ctx,
		"FOR s IN user_sessions FILTER s.session_token == @token FOR u in users FILTER u._key == s.user FOR r in roles FILTER r._key == u.role RETURN merge(s, {user: merge(u, {role: r})})",
		map[string]interface{}{"token": sessionToken})
	if err != nil {
		slog.Error(err)
		return models.SessionToken{}, config.ErrSessionNotFound
	}
	defer result.Close()

	_, err = result.ReadDocument(arango.Ctx, &sessionObject)
	if driver.IsNoMoreDocuments(err) || err != nil {
		slog.Errorf("Failed to read document: %v", err)
		return models.SessionToken{}, config.ErrSessionNotFound
	}

	utils.SessionToSafe(&sessionObject)
	return sessionObject, nil
}
