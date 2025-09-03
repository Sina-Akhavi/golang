// utils/otp.go
package utils

import (
    "crypto/rand"
    "fmt"
    "math/big"
    "otp-service/config"
	"time" // Import the time package
)

func GenerateOTP(phone string) string {
    max := big.NewInt(10000) // OTP range: 0000-9999
    otp, _ := rand.Int(rand.Reader, max)

    otpStr := fmt.Sprintf("%04d", otp.Int64()) // Pad with zeros
	fmt.Println("Generated OTP:", otpStr)     // Debug log
    err := config.RedisClient.Set(config.Ctx, phone, otpStr, 120*time.Second).Err()
	if err != nil {
		fmt.Println("Error storing OTP in Redis:", err) // Debug log
	} else {
		fmt.Println("Successfully stored OTP in Redis for phone:", phone) // Debug log
	}

    fmt.Printf("Generated OTP for %s: %s\n", phone, otpStr) // Print OTP to console
    return otpStr
}

func ValidateOTP(phone, otp string) bool {
    storedOtp, err := config.RedisClient.Get(config.Ctx, phone).Result()
    if err != nil || storedOtp != otp {
        return false
    }

    config.RedisClient.Del(config.Ctx, phone) // Remove OTP after validation
    return true
}