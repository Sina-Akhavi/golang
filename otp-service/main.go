// main.go
package main

import (
    "otp-service/config"
    "otp-service/models"
    "otp-service/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    config.InitRedis()
    models.InitUsers()

    router := gin.Default()
    
    router.POST("/request-otp", handlers.RequestOTP) // Request OTP
    router.POST("/validate-otp", handlers.ValidateOTP) // Validate OTP and login/register
    router.GET("/user", handlers.GetSingleUserByPhone)
    router.GET("/users", handlers.GetUsersWithPagination) // Get paginated list of users


    router.Run(":8080") // Start server on port 8080
}

