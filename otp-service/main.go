// main.go
package main

import (
    "otp-service/config"
    "otp-service/handlers"

    "github.com/gin-gonic/gin"
)

func main() {
    config.InitRedis()

    router := gin.Default()
    router.POST("/request-otp", handlers.RequestOTP) // Request OTP
    router.POST("/validate-otp", handlers.ValidateOTP) // Validate OTP and login/register

    router.Run(":8080") // Start server on port 8080
}