package services

import (
	"github.com/gookit/slog"
	"github.com/yzaimoglu/flathunter/config"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(id string) (models.User, error) {
	mongo := config.NewMongoClient()
	defer mongo.Close()

	var user models.User
	var objId primitive.ObjectID

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		slog.Error(err)
		return models.User{}, config.ErrUserNotFound
	}

	// var filter []bson.M
	// filter = append(filter, bson.M{"_id": objId})
	// then use bson.M{"$and": filter} or bson.M{"$or": filter} in Find

	err = mongo.Client.Database(config.GetString("DB_DATABASE")).Collection("users").FindOne(mongo.Ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		slog.Error(err)
		return models.User{}, config.ErrUserNotFound
	}
	return user, nil
}

func InsertUser(createUser models.CreateUser) (interface{}, error) {
	mongo := config.NewMongoClient()
	defer mongo.Close()

	user := models.User{
		Email:          createUser.Email,
		HashedPassword: createUser.Password,
	}

	result, err := mongo.Client.Database(config.GetString("DB_DATABASE")).Collection("users").InsertOne(mongo.Ctx, user)
	if err != nil {
		slog.Error(err)
		return nil, config.ErrUserInsertError
	}
	return result.InsertedID, nil
}
