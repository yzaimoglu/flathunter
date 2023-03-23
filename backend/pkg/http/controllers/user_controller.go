package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"github.com/yzaimoglu/flathunter/pkg/services"
)

// GetUsers retrieves all the users
func GetUsers(c *fiber.Ctx) error {
	users, err := services.GetUsers()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// GetUser retrieves all the users
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := services.GetUser(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

// GetUsers retrieves all the users
func InsertUser(c *fiber.Ctx) error {
	insertedId, err := services.InsertUser(models.CreateUser{
		Email:    "yagi@mitocho.io",
		Password: "123456",
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"insertedId": insertedId,
	})
}
