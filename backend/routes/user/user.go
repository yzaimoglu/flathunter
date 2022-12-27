package routes

import "github.com/gin-gonic/gin"

// Setup the user routes for the API
func GetUserRoutes(router *gin.RouterGroup) {
	userRoutes := router.Group("/user")
	{
	  userRoutes.GET("/")
  }
}
