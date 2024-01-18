package main

import (
	"net/http"

	"github.com/amartin3659/HttpServerPractice/internal/handlers"
)

func routes() http.Handler {
  
  app.Mux.HandleFunc("/", handlers.Home)
  app.Mux.HandleFunc("/user/login", handlers.Login)
  app.Mux.HandleFunc("/user/profile", handlers.Profile)

  return app.Mux

}
