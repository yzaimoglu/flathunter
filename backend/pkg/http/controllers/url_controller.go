package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"github.com/yzaimoglu/flathunter/pkg/services"
)

// GetURLs retrieves all the urls
func GetURLs(c *fiber.Ctx) error {
	urls, err := services.GetURLs()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(urls)
}

// GetURL retrieves a single url by id
func GetURL(c *fiber.Ctx) error {
	id := c.Params("id")
	url, err := services.GetURL(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(url)
}

// GetUserURLs retrieves all the urls of a user
func GetUserURLs(c *fiber.Ctx) error {
	userId := c.Params("userId")
	urls, err := services.GetUserURLs(userId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(urls)
}

// GetUserURL retrieves all the urls of a user
func GetUserURL(c *fiber.Ctx) error {
	userId := c.Params("userId")
	urlId := c.Params("urlId")
	url, err := services.GetUserURL(userId, urlId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(url)
}

// InsertUserURL inserts a new url for a user
func InsertUserURL(c *fiber.Ctx) error {
	var createUserURL models.CreateUserURLRequest
	if err := c.BodyParser(&createUserURL); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	insertedId, err := services.InsertUserURL(createUserURL)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"insertedId": insertedId,
	})
}
