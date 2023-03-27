package config

import (
	"os"
	"strconv"

	"github.com/gookit/slog"
	"github.com/joho/godotenv"
)

// Load loads the environment variables from the .env file
func Load() {
	err := godotenv.Load()
	if err != nil {
		slog.Errorf("Error loading .env file: %v", err)
	}
}

// GetString returns a string value from the environment
func GetString(key string) string {
	return os.Getenv(key)
}

// GetInteger returns an integer value from the environment
func GetInteger(key string) int {
	int_value, err := strconv.ParseInt(os.Getenv(key), 10, 64)
	if err != nil {
		slog.Error("Could not parse the integer value for key: " + key)
		return 0
	}

	return int(int_value)
}

// GetBoolean returns a boolean value from the environment
func GetBoolean(key string) bool {
	bool_value, err := strconv.ParseBool(os.Getenv(key))
	if err != nil {
		slog.Error("Could not parse the boolean value for key: " + key)
		return false
	}

	return bool_value
}
