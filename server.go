package main

import (
	"log"
	"os"
	"net/http"
	"workspace/controller"
	"workspace/logging"
)

type Server struct {
	userController controller.UserController
}

func NewServer(userController controller.UserController) *Server {
	return &Server{
		userController: userController,
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/signup", s.userController.SignUp)
	
	return logging.Logger(os.Stderr, mux)
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: s.Handler(),
	}
	log.Println("Application starting in port :%s ...", os.Getenv("APP_PORT"))
	
	log.Fatal(httpServer.ListenAndServe())
	
}