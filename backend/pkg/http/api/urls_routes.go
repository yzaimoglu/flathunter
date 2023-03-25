package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yzaimoglu/flathunter/pkg/http/controllers"
)

func urlsRoutes(router fiber.Router) {
	url := router.Group("/urls")
	url.Get("/get/:id", controllers.GetURL)

	userURLs := router.Group("/user_urls")
	userURLs.Get("/get/:userId", controllers.GetUserURLs)
	userURLs.Get("/get/:userId/:urlId", controllers.GetUserURL)
	userURLs.Post("/insert", controllers.InsertUserURL)
	//userURLs.Delete("/delete", controllers.DeleteUserURL)
}
