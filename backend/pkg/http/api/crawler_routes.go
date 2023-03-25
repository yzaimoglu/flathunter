package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yzaimoglu/flathunter/pkg/http/controllers"
)

// crawlerRoutes retrieves all the routes for the crawler
func crawlerRoutes(router fiber.Router) {
	user := router.Group("/crawler")
	user.Get("/urls/get", controllers.GetURLs)
}
