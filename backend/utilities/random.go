package utilities

import (
	"crypto/rand"

	"github.com/google/uuid"
)

// Generate a random UUID for the session
func GenerateRandomSession() []byte {
	uuid := uuid.New()
	return []byte(uuid.String())
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
