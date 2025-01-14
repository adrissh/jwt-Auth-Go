package main

import (
	"jwtAuthGo/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
	port := os.Getenv("PORT")

	router := gin.Default()
	routes.AuthRoute(router)

	if err := router.Run(":" + port); err != nil {
		panic(err)
	}
}
