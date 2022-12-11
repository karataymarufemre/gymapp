package main

import (
	"log"
	"net/http"
	"os"
	"workspace/internal/api/handler"
	"workspace/internal/middleware"
)

type Server struct {
	authHandler   handler.AuthHandler
	jwtMiddleware middleware.JWTMiddleware
}

func NewServer(authHandler handler.AuthHandler, jwtMiddleware middleware.JWTMiddleware) *Server {
	return &Server{
		authHandler:   authHandler,
		jwtMiddleware: jwtMiddleware,
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	s.authHandler.Handle(mux)
	return middleware.Logger(os.Stderr, s.jwtMiddleware.Middleware(mux))
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: s.Handler(),
	}
	log.Println("Application starting in port: ", os.Getenv("APP_PORT"))

	log.Fatal(httpServer.ListenAndServe())

}
