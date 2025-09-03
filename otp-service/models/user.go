package models

import ("time"
		"strconv"
)

// User struct represents a user with phone number and registration date
type User struct {
    Phone           string    `json:"phone"`           // User's phone number
    Name            string    `json:"name"`            // User's name
    RegistrationDate time.Time `json:"registration_date"` // User's registration date
}

// In-memory user store
var Users = map[string]User{} // Map to store users

func InitUsers() {
    for i := 1; i <= 20; i++ {
        phone := "123456" + formatNumber(i) // Generate phone numbers dynamically
        name := "User " + formatNumber(i)  // Generate user names dynamically
        Users[phone] = User{
            Phone:           phone,
            Name:            name,
            RegistrationDate: time.Now(), // Set registration date to current time
        }
    }
}

// Helper function to format numbers as strings with leading zeros (if needed)
func formatNumber(n int) string {
    if n < 10 {
        return "0" + strconv.Itoa(n) // Add leading zero for single digits
    }
    return strconv.Itoa(n)
}