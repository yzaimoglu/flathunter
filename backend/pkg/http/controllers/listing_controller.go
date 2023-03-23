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
