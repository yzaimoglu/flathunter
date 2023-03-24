package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yzaimoglu/flathunter/pkg/http/controllers"
)

func urlsRoutes(router fiber.Router) {
	listing := router.Group("/urls")
	listing.Get("/get", controllers.GetURLs)
	listing.Get("/get/:id", controllers.GetURL)
}
