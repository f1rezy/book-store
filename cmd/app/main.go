package main

import (
	"log"

	"github.com/f1rezy/book-store/internal/config"
	"github.com/f1rezy/book-store/internal/models"
	"github.com/f1rezy/book-store/internal/routes"
	"github.com/f1rezy/book-store/pkg/db/postgresql"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := postgresql.Config{
		Host:     config.EnvString("DB_HOST", ""),
		Port:     config.EnvString("DB_PORT", ""),
		User:     config.EnvString("DB_USER", ""),
		Password: config.EnvString("DB_PASSWORD", ""),
		DBName:   config.EnvString("DB_NAME", ""),
		SSLMode:  config.EnvString("DB_SSLMODE", ""),
	}

	postgresql.InitDB(config)
	if err := postgresql.DB.AutoMigrate(&models.User{}, &models.Book{}); err != nil {
		panic(err)
	}

	routes.AuthRoutes(r)
	routes.BookRoutes(r)

	r.Run(":8080")
}
