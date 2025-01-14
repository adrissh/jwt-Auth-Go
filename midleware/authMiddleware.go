package midleware

import (
	"fmt"
	"jwtAuthGo/response"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	// Retrieve the token from the cookie
	tokenString, err := c.Cookie("access_token")
	if err != nil {
		fmt.Println("Token missing in cookie")
		// c.String(http.StatusForbidden, "you don't have credential")
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Status:     "error",
			StatusCode: http.StatusUnauthorized,
			Message:    "you dont have credentials",
			Timestamp:  time.Now().Format(time.RFC3339),
			Errors:     "Token missing in cookie",
		})
		c.Abort()
		return
	}

	// Parse dan validasi token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validasi algoritma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		// log.Fatal(err)
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Status:     "error",
			StatusCode: http.StatusUnauthorized,
			Message:    "invalid credentials",
			Timestamp:  time.Now().Format(time.RFC3339),
			Errors:     err.Error(),
		})
		c.Abort()
		return
	}

	// get claim token
	claims, _ := token.Claims.(jwt.MapClaims)
	username := claims["username"]

	c.Set("data", username)
	c.Next()

}
