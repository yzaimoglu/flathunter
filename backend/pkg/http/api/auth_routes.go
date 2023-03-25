package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/yzaimoglu/flathunter/pkg/http/controllers"
)

// authRoutes retrieves all the authentication routes
func authRoutes(router fiber.Router) {
	user := router.Group("/auth")
	user.Post("/login", controllers.LoginUser)
	user.Post("/register", controllers.RegisterUser)
	user.Delete("/logout", controllers.LogoutUser)
	user.Get("/session", controllers.UserSession)
}
