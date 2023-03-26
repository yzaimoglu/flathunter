package controllers

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"github.com/yzaimoglu/flathunter/pkg/services"
)

// GetUsers retrieves all the users
func GetUsers(c *fiber.Ctx) error {
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}
	users, err := services.GetUsers(pageInt)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// GetUserByID retrieves a user by the id
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := services.GetUserByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// GetUserByEmail retrieves a user by the email
func GetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	user, err := services.GetUserByEmail(email)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// RegisterUser registers the user
func RegisterUser(c *fiber.Ctx) error {
	var createUser models.CreateUserRequest
	if err := c.BodyParser(&createUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	insertedId, err := services.InsertUser(createUser)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"insertedId": insertedId,
	})
}

// LoginUser logs in the user
func LoginUser(c *fiber.Ctx) error {
	var loginUser models.LoginUserRequest
	if err := c.BodyParser(&loginUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := services.LoginUser(loginUser)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    token.SessionToken,
		HTTPOnly: true,
		MaxAge:   60 * 60 * 24 * 2,
		Path:     "/",
		Secure:   true,
	})

	return c.Status(fiber.StatusOK).JSON(token)
}

// LogoutUser logs out the user
func LogoutUser(c *fiber.Ctx) error {
	sessionTokenCookie := strings.ReplaceAll(c.Cookies("session_token"), "session_token=", "")
	if sessionTokenCookie == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "session token not found",
		})
	}

	err := services.LogoutUser(sessionTokenCookie)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.ClearCookie("session_token")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}

// UserSession retrieves the user session
func UserSession(c *fiber.Ctx) error {
	sessionTokenCookie := strings.ReplaceAll(c.Cookies("session_token"), "session_token=", "")
	if sessionTokenCookie == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "session token not found",
		})
	}

	sessionToken, err := services.GetSessionWithUserByToken(sessionTokenCookie)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(sessionToken)
}

// ChangePassword changes the password
func ChangePassword(c *fiber.Ctx) error {
	var changePasswordRequest models.ChangePasswordRequest
	if err := c.BodyParser(&changePasswordRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err := services.ChangePassword(changePasswordRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password changed successfully",
	})
}
