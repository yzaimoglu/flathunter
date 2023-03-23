package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/gookit/slog"
)

const (
	ArangoConnectionString  = "http://%s:%d"
	ArangoConnectionTimeout = 5
)

// ArangoConnection is a struct that holds the connection information for the ArangoDB database.
type ArangoConnection struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// ArangoClient is a struct that holds the client and context for the ArangoDB database.
type ArangoClient struct {
	Connection ArangoConnection
	Client     driver.Client
	Ctx        context.Context
	Cancel     context.CancelFunc
	Database   driver.Database
}

// NewArangoClient creates a new ArangoDB client.
func NewArangoClient() *ArangoClient {
	arangoConnection := ArangoConnection{
		Host:     GetString("DB_HOST"),
		Port:     GetInteger("DB_PORT"),
		Database: GetString("DB_DATABASE"),
		Username: GetString("DB_USER"),
		Password: GetString("DB_PASSWORD"),
	}

	connectionURI := fmt.Sprintf(ArangoConnectionString, arangoConnection.Host, arangoConnection.Port)

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{connectionURI},
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
		ConnLimit: 100,
	})
	if err != nil {
		slog.Fatalf("Failed to create connection: %v", err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(arangoConnection.Username, arangoConnection.Password),
	})
	if err != nil {
		slog.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), ArangoConnectionTimeout*time.Second)

	arangoClient := &ArangoClient{
		Client:     client,
		Ctx:        ctx,
		Cancel:     cancel,
		Connection: arangoConnection,
	}

	return arangoClient.GetDatabase()
}

// SetupArango creates the database if it does not exist.
func SetupArango() {
	slog.Info("Setting up the ArangoDB database...")
	arango := NewArangoClient()
	defer arango.Cancel()

	if !arango.CheckDatabase() {
		arango.CreateDatabase()
	}

	arango.CheckCollectionsAndCreate()
	slog.Info("Setup of the ArangoDB database complete.")
}

// CreateDatabase creates the flathunter database.
func (arango *ArangoClient) CreateDatabase() (driver.Database, error) {
	db, err := arango.Client.CreateDatabase(arango.Ctx, arango.Connection.Database, &driver.CreateDatabaseOptions{})
	if err != nil {
		slog.Errorf("Failed to create database: %v", err)
		return nil, err
	}

	arango.Database = db
	return db, nil
}

// CheckDatabase checks if the flathunter database exists.
func (arango *ArangoClient) CheckDatabase() bool {
	exists, err := arango.Client.DatabaseExists(arango.Ctx, arango.Connection.Database)
	if err != nil {
		slog.Fatalf("Failed to check if database exists: %v", err)
	}
	return exists
}

// CheckCollectionsAndCreate checks if the collections exist and creates them if they do not.
func (arango *ArangoClient) CheckCollectionsAndCreate() {
	if !arango.CheckCollection("users") {
		arango.CreateCollection("users")
	}
	if !arango.CheckCollection("platforms") {
		arango.CreateCollection("platforms")
	}
	if !arango.CheckCollection("roles") {
		arango.CreateCollection("roles")
	}
	if !arango.CheckCollection("urls") {
		arango.CreateCollection("urls")
	}
	if !arango.CheckCollection("listings") {
		arango.CreateCollection("listings")
	}
}

// CheckCollection checks if a collection exists.
func (arango *ArangoClient) CheckCollection(name string) bool {
	collections, err := arango.Database.Collections(arango.Ctx)
	if err != nil {
		slog.Fatalf("Failed to check if collections exist: %v", err)
	}

	for _, collection := range collections {
		if collection.Name() == name {
			return true
		}
	}

	return false
}

// CreateCollection creates a collection.
func (arango *ArangoClient) CreateCollection(name string) error {
	if arango.CheckCollection(name) {
		slog.Infof("Collection %s already exists.", name)
		return nil
	}

	if _, err := arango.Database.CreateCollection(arango.Ctx, name, nil); err != nil {
		slog.Errorf("Failed to create collection: %v", err)
		return err
	}
	return nil
}

// GetDatabase retrieves the flathunter database.
func (arango *ArangoClient) GetDatabase() *ArangoClient {
	db, err := arango.Client.Database(arango.Ctx, arango.Connection.Database)
	if err != nil {
		slog.Fatalf("Failed to retrieve database: %v", err)
	}

	arango.Database = db
	return arango
}

// Close closes the ArangoDB client.
func (arango *ArangoClient) Close() {
	arango.Cancel()
}
