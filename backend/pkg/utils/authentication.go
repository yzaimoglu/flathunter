package utils

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hash a plain password with the bcrypt library
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

// CheckPassword check a hashed password with a plain password
func CheckPassword(hashedPassword string, plainPassword string) bool {
	byteHash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPassword))

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// CreateSession creates a new session for the user
func CreateSession(user models.User) (models.SessionToken, error) {
	sessionToken := models.SessionToken{
		User:         user,
		CreatedAt:    time.Now().Unix(),
		ExpiresAt:    time.Now().Add(time.Hour * 48).Unix(),
		SessionToken: ToBase64([]byte(GenerateRandomSession())),
	}
	return sessionToken, nil
}

// Generate a random UUID for the session
func GenerateRandomSession() string {
	first_uuid := SHA512(uuid.New().String())
	second_uuid := SHA512(uuid.New().String())
	return first_uuid + second_uuid
}

func RemoveSession(c *fiber.Ctx) {
	//session_token := c.Cookies("session_token")

}
