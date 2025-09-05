package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"otp-service/utils" // for jwtKey and Claim type

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RequireToken() gin.HandlerFunc {
  return func(c *gin.Context) {
    fmt.Println("Hi Ali!!!")
    auth := c.GetHeader("Authorization")
    fmt.Println("AUTH=", auth)
    if auth == "" {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
      return
    }

    parts := strings.Fields(auth)
    if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization format must be Bearer <token>"})
      return
    }

    tokenStr := parts[1]
    token, err := jwt.ParseWithClaims(tokenStr, &utils.Claims{}, func(t *jwt.Token) (interface{}, error) {
      return utils.JwtKey, nil
    })
    if err != nil || !token.Valid {
      c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
      return
    }

    // // optionally pull the phone number out of the claims and store it in context
    // if claims, ok := token.Claims.(*utils.Claims); ok {
    //   c.Set("phone", claims.Phone)
    // }

    c.Next()
  }
}