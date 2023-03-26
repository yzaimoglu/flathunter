package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yzaimoglu/flathunter/pkg/http/controllers"
)

func listingRoutes(router fiber.Router) {
	listing := router.Group("/listings")
	listing.Get("/get", controllers.GetListings)
	listing.Get("/get/:listingId", controllers.GetListing)

	userListings := router.Group("/user_listings")
	userListings.Get("/get/:userId", controllers.GetUserListings)
	userListings.Get("/get/:userId/:listingId", controllers.GetUserListing)
	userListings.Post("/insert", controllers.InsertUserListing)
	userListings.Delete("/delete/:listingId", controllers.DeleteUserListing)
}
