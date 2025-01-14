package services

import (
	"fmt"
	"jwtAuthGo/config"
	"jwtAuthGo/models"
	"jwtAuthGo/payloads"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type CustomClaims struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func UserAuthentication(reqPayloads payloads.AuthPayload) (models.User, error) {

	config.ConnectDB()
	var dbUser models.User
	query := config.DB.Where("username = ?", reqPayloads.Username).First(&dbUser).Error
	if query != nil {
		fmt.Println("error bro : ", query)
		return models.User{}, query
		// return
	}
	log.Println("Id User :", dbUser.ID)

	// verify passwd
	verifyPasswd := bcrypt.CompareHashAndPassword([]byte(dbUser.Password_hash), []byte(reqPayloads.Password))
	if verifyPasswd != nil {
		fmt.Println("passwd incorrect")
		return models.User{}, verifyPasswd
		// return
	}
	// fmt.Println("login succesfully")

	// GenerateAccessToken(dbUser)
	// GenerateRefreshToken(dbUser)

	return dbUser, nil
}

func GenerateAccessToken(data models.User) (token string, err error) {
	// generate accsess token
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	claims := CustomClaims{
		Id:       data.ID,
		Username: data.Username,
		Email:    data.Email,
		Role:     data.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 3)),
		},
	}

	// cretae access token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// sign token with secret key
	signedToken, err := accessToken.SignedString(secretKey)
	if err != nil {
		fmt.Println("error when sign token", err)
	}
	return signedToken, nil

}

func GenerateRefreshToken(data models.User) (token string, err error) {
	secretKey := []byte(os.Getenv("SECRET_KEY_REFRESH_TOKEN"))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": data.ID,
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
	})

	// sign token with secret key
	signedToken, err := claims.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
	}
	return signedToken, nil
}

func UserLogout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

}
