package controllers

import "github.com/gofiber/fiber/v2"

func GetListings(c *fiber.Ctx) error {
	return c.SendString("Listing")
}
