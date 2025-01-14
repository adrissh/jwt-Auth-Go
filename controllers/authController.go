package controllers

import (
	"fmt"
	"jwtAuthGo/payloads"
	"jwtAuthGo/response"
	"jwtAuthGo/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HandlerAuth(c *gin.Context) {
	var data payloads.AuthPayload
	if err := c.ShouldBind(&data); err != nil {
		fmt.Println(err)
	}

	result, err := services.UserAuthentication(data)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{
			Status:     "error",
			StatusCode: http.StatusUnauthorized,
			Message:    "invalid credentials",
			Timestamp:  time.Now().Format(time.RFC3339),
			Errors:     "invalid username or password ",
		})
		return
	}

	// set acces token to  cookies
	accessToken, _ := services.GenerateAccessToken(result)
	c.SetCookie("access_token", accessToken, 0, "/", "", true, true)
	// set refresh token to cookies
	refreshToken, _ := services.GenerateRefreshToken(result)
	c.SetCookie("refresh_token", refreshToken, 0, "/", "", true, true)
	c.JSON(http.StatusOK, response.SuccessResponse{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Authentication successfully",
		Timestamp:  time.Now().Format(time.RFC3339),
		Payload:    result,
	})

}

func HandlerLogout(c *gin.Context) {
	services.UserLogout(c)
	c.JSON(http.StatusOK, response.SuccessResponse{
		Status:     "success",
		StatusCode: http.StatusOK,
		Message:    "Succesfully destroy session",
		Timestamp:  time.Now().Format(time.RFC3339),
		Payload:    nil,
	})
}
