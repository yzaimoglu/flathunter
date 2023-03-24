package controllers

import (
	"github.com/gofiber/fiber/v2"
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
