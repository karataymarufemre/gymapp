package main

import (
	"log"
	"net/http"
	"os"
	"workspace/controller"
	"workspace/middleware"
)

type Server struct {
	userController controller.UserController
	jwtMiddleware  middleware.JWTMiddleware
}

func NewServer(userController controller.UserController, jwtMiddleware middleware.JWTMiddleware) *Server {
	return &Server{
		userController: userController,
		jwtMiddleware:  jwtMiddleware,
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/signup", s.userController.SignUp)
	mux.HandleFunc("/signin", s.userController.SignIn)
	mux.Handle("/refresh-token", s.jwtMiddleware.MiddlewareForAccessToken(http.HandlerFunc(s.userController.GetAccessToken)))
	return middleware.Logger(os.Stderr, mux)
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: s.Handler(),
	}
	log.Println("Application starting in port: ", os.Getenv("APP_PORT"))

	log.Fatal(httpServer.ListenAndServe())

}
