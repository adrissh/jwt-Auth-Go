package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := os.Getenv("HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	passwd := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	// DSN (Data Source Name) for connection to PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, passwd, db_name, port)

	// open connection to postgresql
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Jika koneksi gagal, log error dan hentikan program
		log.Fatalf("Connection failed: %v", err)
	}

	// Jika koneksi berhasil
	fmt.Println("Connection successfully")
}
