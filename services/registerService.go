package services

import (
	"errors"
	"fmt"
	"jwtAuthGo/config"
	"jwtAuthGo/models"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func HashPasswd(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err

}

func GenerateUuid() (string, error) {
	id, err := gonanoid.New()
	return id, err
}

func UserRegister(user models.User) []string {
	// start transaction
	config.ConnectDB()

	var errorCollection []string

	checkExists := func(field, value string) bool {
		err := config.DB.Where(field+"= ?", value).First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return false
			}
			fmt.Println("error when checking")
		}
		return true
	}

	if checkExists("username", user.Username) {
		errorCollection = append(errorCollection, "username already exists")
	}
	if checkExists("email", user.Email) {
		errorCollection = append(errorCollection, "email already exists")
	}
	if len(errorCollection) > 0 {
		return errorCollection
	}

	fmt.Println("all field not found , next to store to DB")
	result := config.DB.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return nil

}
