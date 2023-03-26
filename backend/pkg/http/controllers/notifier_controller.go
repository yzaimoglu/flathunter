package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"github.com/yzaimoglu/flathunter/pkg/services"
)

// GetNotifiers returns all notifiers for a user
func GetNotifiers(c *fiber.Ctx) error {
	userId := c.Params("userId")
	notifiers, err := services.GetNotifiers(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(notifiers)
}

// GetNotifier returns a notifier for a user
func GetNotifier(c *fiber.Ctx) error {
	userId := c.Params("userId")
	notifierId := c.Params("notifierId")
	notifier, err := services.GetNotifier(userId, notifierId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(notifier)
}

// InsertNotifier inserts a new notifier for a user
func InsertNotifier(c *fiber.Ctx) error {
	var createNotifier models.CreateNotifier

	if err := c.BodyParser(&createNotifier); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	id, err := services.InsertNotifier(models.CreateNotifier{
		User:    createNotifier.User,
		Type:    createNotifier.Type,
		Options: createNotifier.Options,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"insertedId": id,
	})
}

// DeleteNotifier deletes a notifier for a user
func DeleteNotifier(c *fiber.Ctx) error {
	userId := c.Params("id")
	listingId := c.Params("listingId")

	err := services.DeleteNotifier(userId, listingId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"deletedId": listingId,
	})
}
