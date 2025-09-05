// main.go
package main

import (
	"otp-service/config"
	"otp-service/handlers"
	"otp-service/middlewares"
	"otp-service/models"

	_ "otp-service/docs" // Import the generated Swagger docs

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // Swagger files handler
	"github.com/swaggo/gin-swagger" // Swagger UI integration
)

func main() {
    config.InitRedis()
    models.InitUsers()

    router := gin.Default()
    
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.POST("/request-otp", handlers.RequestOTP) 
    router.POST("/validate-otp", handlers.ValidateOTP)

    auth := router.Group("/") // Create a group for protected routes
    auth.Use(middlewares.RequireToken()) // Apply the RequireToken middleware
    {
        auth.GET("/user", handlers.GetSingleUserByPhone)
        auth.GET("/users", handlers.GetUsersWithPagination) // Get paginated list of users
        auth.POST("/users", handlers.CreateUser)            // Create new user
    }
    router.Run(":8080") // Start server on port 8080
}
