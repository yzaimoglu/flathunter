package services

import (
	"github.com/arangodb/go-driver"
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/pkg/config"
	"github.com/yzaimoglu/flathunter/pkg/models"
)

// GetNotifiers returns all notifiers for a user
func GetNotifiers(userId string) ([]models.Notifier, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var notifiers []models.Notifier

	result, err := arango.Database.Query(arango.Ctx,
		"FOR notifier IN user_notifiers FILTER notifier.user == @id FOR user IN users FILTER user._key == notifier.user FOR role IN roles FILTER role._key == user.role RETURN merge(notifier, {user: merge(user, {role: role})})",
		map[string]interface{}{"id": userId})
	if err != nil {
		slog.Error(err)
		return []models.Notifier{}, err
	}
	defer result.Close()

	for {
		var notifier models.Notifier
		_, err := result.ReadDocument(arango.Ctx, &notifier)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			slog.Errorf("Failed to read document: %v", err)
			return nil, err
		}
		notifiers = append(notifiers, notifier)
	}

	if len(notifiers) == 0 {
		return []models.Notifier{}, config.ErrNotifierNotFound
	}

	return notifiers, nil
}

// GetNotifier returns a notifier for a user
func GetNotifier(userId string, listingId string) (models.Notifier, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	var notifier models.Notifier

	result, err := arango.Database.Query(arango.Ctx,
		"FOR notifier IN user_notifiers FILTER notifier.user == @userId && notifier._key == @listingId FOR user IN users FILTER user._key == notifier.user FOR role IN roles FILTER role._key == user.role RETURN merge(notifier, {user: merge(user, {role: role})})",
		map[string]interface{}{"userId": userId, "listingId": listingId})
	if err != nil {
		slog.Error(err)
		return models.Notifier{}, err
	}
	defer result.Close()

	_, err = result.ReadDocument(arango.Ctx, &notifier)
	if err != nil {
		slog.Errorf("Failed to read document: %v", err)
		return models.Notifier{}, config.ErrNotifierNotFound
	}

	return notifier, nil
}

// InsertNotifier inserts a new notifier for a user
func InsertNotifier(createNotifier models.CreateNotifier) (string, error) {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUserNotifiersCollection)
	if err != nil {
		slog.Errorf("Failed to get collection %s: %v", config.ArangoUserNotifiersCollection, err)
		return "", err
	}

	meta, err := collection.CreateDocument(arango.Ctx, models.CreateNotifier{
		User:    createNotifier.User,
		Type:    createNotifier.Type,
		Options: createNotifier.Options,
	})
	if err != nil {
		slog.Errorf("Failed to create document: %v", err)
		return "", err
	}

	return meta.Key, nil
}

// DeleteNotifier deletes a notifier for a user
func DeleteNotifier(userId string, notifierId string) error {
	arango := config.NewArangoClient()
	defer arango.Close()

	collection, err := arango.Database.Collection(arango.Ctx, config.ArangoUserNotifiersCollection)
	if err != nil {
		slog.Errorf("Failed to get collection %s: %v", config.ArangoUserNotifiersCollection, err)
		return err
	}

	_, err = collection.RemoveDocument(arango.Ctx, notifierId)
	if err != nil {
		slog.Errorf("Failed to remove document: %v", err)
		return err
	}

	return nil
}
