package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"workspace/internal/api/controller"
	"workspace/internal/api/handler"
	"workspace/internal/middleware"
	"workspace/internal/repository"
	"workspace/internal/service"
	"workspace/internal/service/authservice"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	//load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//database connection
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	}
	//dependency injection, initialize services

	//----- user
	authService := authservice.NewAuthService([]byte(os.Getenv("JWT_KEY")))
	userRepository := repository.NewUserRepository(db, authService)

	userService := service.NewUserService(userRepository, authService)
	userController := controller.NewUserController(userService)
	jwtMiddleware := middleware.NewJWTMiddleware(authService)
	authHandler := handler.NewAuthHandler(userController)
	//server
	server := NewServer(authHandler, jwtMiddleware)
	server.Run()
}
