package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yzaimoglu/flathunter/pkg/services"
)

// GetUserListing retrieves a user listing of a user
func GetUserListing(c *fiber.Ctx) error {
	userId := c.Params("userId")
	listingId := c.Params("listingId")
	listing, err := services.GetUserListing(userId, listingId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(listing)
}
