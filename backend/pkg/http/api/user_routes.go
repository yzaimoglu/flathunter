package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yzaimoglu/flathunter/pkg/http/controllers"
)

// userRoutes retrieves all the user routes
func userRoutes(router fiber.Router) {
	user := router.Group("/user")
	user.Get("/get_id/:id", controllers.GetUserByID)
	user.Get("/get_email/:email", controllers.GetUserByEmail)
	user.Put("/change_password", controllers.ChangePassword)

	users := router.Group("/users")
	users.Get("/get", controllers.GetUsers)
}
