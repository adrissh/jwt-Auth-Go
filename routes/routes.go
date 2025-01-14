package routes

import (
	"fmt"
	"jwtAuthGo/controllers"
	"jwtAuthGo/midleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRoute(router *gin.Engine) {
	v1 := router.Group("/api/v1.0")
	v1.POST("register", midleware.ValidateRegister, controllers.HandleRegister)
	v1.POST("login", controllers.HandlerAuth)
	v1.POST("logout", controllers.HandlerLogout)
	v1.GET("dashboard", midleware.AuthMiddleware, func(c *gin.Context) {
		fmt.Println("Welcome dashboard")
		myName := c.MustGet("data").(string)
		reponseMessage := fmt.Sprintf("welcome to dasboard %s\n", myName)
		c.String(http.StatusOK, reponseMessage)

	})

}
