package database

import (
	"fmt"
	"jwtAuthGo/config"
	"jwtAuthGo/models"
	"log"
)

func MigrateTable() {
	config.ConnectDB()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error during migration :", err)
		return

	}
	fmt.Println("migrate success ...")
}
