package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yzaimoglu/flathunter/pkg/http/controllers"
)

// notifierRoutes retrieves all the routes for the notifier
func notifierRoutes(router fiber.Router) {
	notifier := router.Group("/notifier")
	notifier.Get("/get/:userId", controllers.GetNotifiers)
	notifier.Get("/get/:userId/:notifierId", controllers.GetNotifier)
	notifier.Post("/insert", controllers.InsertNotifier)
	notifier.Delete("/delete/:userId/:listingId", controllers.DeleteNotifier)
}
