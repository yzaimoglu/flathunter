package config

import (
  "os"
  "time"
)

// Get a specific environment variable
func GetEnv(key, defaultValue string) string {
  if value, ok := os.LookupEnv(key); ok {
    return value
  }
  return defaultValue
}

// Get the current time as a UNIX Timestamp
func GetCurrentTime() int64 {
  return time.Now().UnixNano() / int64(time.Millisecond)
}
