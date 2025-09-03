// utils/jwt.go
package utils

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

var jwtKey = []byte("secret_key") // Replace with a secure key

type Claims struct {
    Phone string `json:"phone"`
    jwt.StandardClaims
}

func GenerateJWT(phone string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours
    claims := &Claims{
        Phone: phone,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}