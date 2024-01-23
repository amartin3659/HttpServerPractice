package main

import (
	"net/http"

	"github.com/amartin3659/HttpServerPractice/internal/handlers"
)

func routes() http.Handler {

	app.Mux.HandleFunc("/home", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
      handlers.Repo.Home(w, r)
    default:
      http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
      return
    }
  }))
	app.Mux.HandleFunc("/user/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.Repo.GetLogin(w, r)
		case http.MethodPost:
			handlers.Repo.PostLogin(w, r)
		}
	}))
	app.Mux.HandleFunc("/error", handlers.Repo.ErrorPage)
	app.Mux.HandleFunc("/user/logout", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.Repo.GetLogout(w, r)
		case http.MethodPost:
			handlers.Repo.PostLogout(w, r)
		}
	}))
	app.Mux.HandleFunc("/user/profile/", handlers.Repo.Profile)
	app.Mux.HandleFunc("/post/", handlers.Repo.GetPost)
	app.Mux.HandleFunc("/user/post", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.Repo.GetNewPost(w, r)
		case http.MethodPost:
			handlers.Repo.PostNewPost(w, r)
		}
	}))
	postPath := "/user/post/"
	app.Mux.Handle(postPath, http.StripPrefix(postPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  	switch r.Method {
	  case http.MethodGet:
		  handlers.Repo.GetUpdatePost(w, r)
	  case http.MethodPost:
		  handlers.Repo.PostUpdatePost(w, r)
	  }
  })))

	return app.Mux

}
