package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yzaimoglu/flathunter/pkg/http/controllers"
)

func listingRoutes(router fiber.Router) {
	listing := router.Group("/listings")
	listing.Get("/", controllers.GetListings)
}
