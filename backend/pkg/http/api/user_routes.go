package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yzaimoglu/flathunter/pkg/http/controllers"
)

// userRoutes retrieves all the user routes
func userRoutes(router fiber.Router) {
	user := router.Group("/")
	user.Get("/user/get_id/:id", controllers.GetUserByID)
	user.Get("/user/get_email/:email", controllers.GetUserByEmail)
	user.Get("/users/get", controllers.GetUsers)
	user.Post("/user/register", controllers.RegisterUser)
}
