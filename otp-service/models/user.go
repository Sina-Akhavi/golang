// models/user.go
package models

import "strconv"

type User struct {
    Phone string `json:"phone"`
    Name  string `json:"name"`
}

var Users = map[string]User{} // In-memory user store

func InitUsers() {
    // Initialize 20 users for testing
    for i := 1; i <= 20; i++ {
        phone := "123456" + formatNumber(i) // Generate phone numbers dynamically
        name := "User " + formatNumber(i)  // Generate user names dynamically
        Users[phone] = User{
            Phone: phone,
            Name:  name,
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
