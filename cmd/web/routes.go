package main

import (
	"net/http"

	"github.com/amartin3659/HttpServerPractice/internal/handlers"
)

func routes() http.Handler {
  
  app.Mux.HandleFunc("/", handlers.Repo.Home)
  app.Mux.HandleFunc("/user/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      handlers.Repo.GetLogin(w, r)
    case http.MethodPost:
      handlers.Repo.PostLogin(w, r)
    }
  }))
  app.Mux.HandleFunc("/user/logout", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      handlers.Repo.GetLogout(w, r)
    case http.MethodPost:
      handlers.Repo.PostLogout(w, r)
    }
  }))
  app.Mux.HandleFunc("/user/profile/", handlers.Repo.Profile)
  app.Mux.HandleFunc("/post/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      handlers.Repo.GetPost(w, r)
    case http.MethodPost:
      handlers.Repo.PostPost(w, r)
    case http.MethodPut:
      handlers.Repo.UpdatePost(w, r)
    }
  }))

  return app.Mux

}
