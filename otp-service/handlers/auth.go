// handlers/auth.go
package handlers

import (
    "net/http"
    "otp-service/models"
    "otp-service/utils"
    "time"
    "github.com/gin-gonic/gin"
)

var rateLimiter = utils.NewRateLimiter(3, 10*time.Minute) // 3 requests per 10 minutes


func RequestOTP(c *gin.Context) {
    var request struct {
        Phone string `json:"phone"`
    }

    err := c.BindJSON(&request) // Parse the JSON and save any error into `err`
    if err != nil {             // Check if there is an error
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"}) // Handle the error
        return                  // Stop the function
    }

    phone := request.Phone
    if !rateLimiter.IsAllowed(phone) {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests. Please wait before trying again."})
		return
	}

    otp := utils.GenerateOTP(request.Phone)
    c.JSON(http.StatusOK, gin.H{"message": "OTP generated", "otp": otp})
}

func ValidateOTP(c *gin.Context) {
    var request struct {
        Phone string `json:"phone"`
        OTP   string `json:"otp"`
    }

    err := c.BindJSON(&request) // Parse the JSON and save any error into `err`
    if err != nil {             // Check if there is an error
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"}) // Handle the error
        return                  // Stop the function
    }

    if !utils.ValidateOTP(request.Phone, request.OTP) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
        return
    }

    // Check if user exists
    user, exists := models.Users[request.Phone]
    if !exists {
        // Register user
        user = models.User{Phone: request.Phone, Name: ""}
        models.Users[request.Phone] = user
    }

    // Generate JWT token
    token, err := utils.GenerateJWT(request.Phone)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Success", "token": token, "user": user})
}