package controllers

import (
	"jwtAuthGo/midleware"
	"jwtAuthGo/models"
	"jwtAuthGo/response"
	"jwtAuthGo/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type DataResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func HandleRegister(c *gin.Context) {
	payload, _ := c.Get("payload")                 // get payload from midleware
	userPayload, ok := payload.(midleware.Request) // Convert payload to middleware.Payload type using type assertion.
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid payload format"})
		return
	}

	// Generate uuid and merger req body
	result, _ := services.HashPasswd(userPayload.Password_hash)
	id, _ := services.GenerateUuid()
	uuid := id
	username := userPayload.Username
	email := userPayload.Email
	password := result
	role := userPayload.Role

	// struct for store to db
	data := models.User{
		Uuid:          uuid,
		Username:      username,
		Email:         email,
		Password_hash: password,
		Role:          role,
	}

	// struct for return response API
	dataResponse := DataResponse{
		Username: username,
		Email:    email,
		Role:     role,
	}

	// database.MigrateTable()
	err := services.UserRegister(data)
	if len(err) > 0 {
		c.JSON(http.StatusConflict, response.ErrorResponse{
			Status:     "error",
			StatusCode: http.StatusConflict,
			Message:    "The field unique already exist in databse",
			Timestamp:  time.Now().Format(time.RFC3339),
			Errors:     err,
		})
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Success register User",
		Timestamp:  time.Now().Format(time.RFC3339),
		Payload:    dataResponse,
	})

}
