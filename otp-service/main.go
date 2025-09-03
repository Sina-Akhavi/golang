// main.go
package main

import (
	"otp-service/config"
	"otp-service/handlers"
	"otp-service/models"

	"github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger" // Swagger UI integration
	"github.com/swaggo/files"       // Swagger files handler
	_ "otp-service/docs"            // Import the generated Swagger docs
)

func main() {
    config.InitRedis()
    models.InitUsers()

    router := gin.Default()
    
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.POST("/request-otp", handlers.RequestOTP) // Request OTP
    router.POST("/validate-otp", handlers.ValidateOTP) // Validate OTP and login/register
    router.GET("/user", handlers.GetSingleUserByPhone)
    router.GET("/users", handlers.GetUsersWithPagination) // Get paginated list of users
    router.POST("/users", handlers.CreateUser)            // Create new user

    router.Run(":8080") // Start server on port 8080
}
