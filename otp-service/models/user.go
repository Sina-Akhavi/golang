// models/user.go
package models

type User struct {
    Phone string `json:"phone"`
    Name  string `json:"name"`
}

var Users = map[string]User{} // In-memory user store