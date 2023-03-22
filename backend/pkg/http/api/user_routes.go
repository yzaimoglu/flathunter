package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yzaimoglu/flathunter/pkg/http/controllers"
)

// userRoutes retrieves all the user routes
func userRoutes(router fiber.Router) {
	user := router.Group("/users")
	user.Get("/get/:id", controllers.GetUser)
	user.Get("/get", controllers.GetUsers)
	user.Get("/insert", controllers.InsertUser)
}
