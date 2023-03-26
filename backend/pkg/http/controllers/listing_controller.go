package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yzaimoglu/flathunter/pkg/models"
	"github.com/yzaimoglu/flathunter/pkg/services"
)

// GetListing retrieves a listing
func GetListing(c *fiber.Ctx) error {
	listingId := c.Params("listingId")
	listing, err := services.GetListing(listingId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(listing)
}

// GetListings retrieves 25 listings
func GetListings(c *fiber.Ctx) error {
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	listings, err := services.GetListings(pageInt)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(listings)
}

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

// GetUserListings retrieves all user listings of a user
func GetUserListings(c *fiber.Ctx) error {
	userId := c.Params("userId")
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	listings, err := services.GetUserListings(userId, pageInt)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(listings)
}

// InsertUserListing inserts a user listing
func InsertUserListing(c *fiber.Ctx) error {
	var createUserListing models.CreateUserListing
	if err := c.BodyParser(&createUserListing); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	_, err := services.InsertUserListing(createUserListing)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

// DeleteUserListing deletes a user listing
func DeleteUserListing(c *fiber.Ctx) error {
	listingId := c.Params("listingId")
	err := services.DeleteUserListing(listingId)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "deleted user listing",
	})
}
