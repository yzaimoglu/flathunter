package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yzaimoglu/flathunter/backend/config"
	"github.com/yzaimoglu/flathunter/backend/middleware"
	routes "github.com/yzaimoglu/flathunter/backend/routes/user"
	models "github.com/yzaimoglu/flathunter/models/auth"
)

func main() {
	// Setup the Database and create the tables
	db := config.SetupDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})

	// Initialize the main router
	gin.SetMode(gin.ReleaseMode)
	mainRouter := gin.New()

	// Add the database to the context
	mainRouter.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Setup default security measures
	mainRouter.Use(middleware.Default())

	// Setup the CORS middleware
	mainRouter.Use(middleware.CORSMiddleware)

	// Setup the Basic and Security Middleware provided by Gin
	mainRouter.Use(gin.Logger())
	mainRouter.Use(gin.Recovery())

	// Standard NoRoute Response
	mainRouter.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "error": "not found"})
	})

	// Standard NoMethod Response
	mainRouter.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"status": http.StatusMethodNotAllowed, "error": "method not allowed"})
	})

	// Set the Favicon
	mainRouter.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// Create the main Route group for the API
	v1 := mainRouter.Group("/v1")
	{
		routes.GetUserRoutes(v1)
	}

	// Run server
	serverPort := fmt.Sprint(config.GetEnv("PORT", fmt.Sprint(8000)))
	fmt.Println("Flathunter REST API started running on port " + serverPort)
	mainRouter.Run(":" + serverPort)
}
