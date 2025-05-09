package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mi759/haikuapi/internal/handler"
	"github.com/mi759/haikuapi/internal/repository"
	"github.com/mi759/haikuapi/internal/usecase"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func buildDSN() string {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	timezone := os.Getenv("TZ")

	return fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=%s", username, password, dbname, timezone)
}

func main() {
	dsn := buildDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
	}))

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	handler.NewUserHandler(router, userUsecase)
	router.Run(":8080")
}
