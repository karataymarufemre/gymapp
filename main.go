package main

import (
	"fmt"
    "log"
	"os"
	"database/sql"
	"workspace/controller"
	"workspace/service"
	"workspace/repository"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
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
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	//server
	server := NewServer(userController)
	server.Run()
}