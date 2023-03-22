package config

import (
	"context"
	"fmt"
	"time"

	"github.com/gookit/slog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mongo constants
const (
	MongoConnectionString  = "mongodb://%s:%s@%s:%d"
	MongoConnectionTimeout = 5
)

// MongoConnection is a struct that holds the connection information for the MongoDB database.
type MongoConnection struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// MongoClient is a struct that holds the client and context for the MongoDB database.
type MongoClient struct {
	Client *mongo.Client
	Ctx    context.Context
	Cancel context.CancelFunc
}

// MongoFlathunter is a struct that holds the MongoDB client and database.
type MongoFlathunter struct {
	MongoClient *MongoClient
	Database    *mongo.Database
}

// NewMongoClient creates a new MongoDB client.
func NewMongoClient() *MongoClient {
	connection := MongoConnection{
		Host:     GetString("DB_HOST"),
		Port:     GetInteger("DB_PORT"),
		Database: GetString("DB_DATABASE"),
		Username: GetString("DB_USER"),
		Password: GetString("DB_PASSWORD"),
	}

	slog.Info("Connecting to the MongoDB database...")

	connectionURI := fmt.Sprintf(MongoConnectionString, connection.Username, connection.Password, connection.Host, connection.Port)

	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		slog.Fatal("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), MongoConnectionTimeout*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		slog.Fatal("Failed to connect to the MongoDB database: %v", err)
	}

	slog.Info("Connected to the MongoDB database.")

	return &MongoClient{
		Client: client,
		Ctx:    ctx,
		Cancel: cancel,
	}
}

// Close closes the MongoDB client.
func (client *MongoClient) Close() {
	slog.Info("Closing the MongoDB client...")
	client.Cancel()
	client.Client.Disconnect(client.Ctx)
	slog.Info("Closed the MongoDB client.")
}
