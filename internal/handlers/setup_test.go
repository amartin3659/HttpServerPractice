package handlers

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/amartin3659/HttpServerPractice/internal/config"
	"github.com/amartin3659/HttpServerPractice/internal/driver"
	"github.com/amartin3659/HttpServerPractice/internal/helpers"
	"github.com/amartin3659/HttpServerPractice/internal/session"
	"github.com/amartin3659/HttpServerPractice/migrations"
)

var app config.AppConfig

func TestMain(m *testing.M) {
  err := run()
  if err != nil {
    fmt.Println("Error")
  }
  
  os.Exit(m.Run())
}

func run() error {
  
  app.SetInProduction(false)
  app.SetInfoLog()
  app.SetErrorLog()
  app.SetMux(http.NewServeMux())
  app.SetSession(session.New())

  db := driver.NewDB()
  seed := migrations.NewSeed(db)
  seed.Seed()
  repo := NewRepo(&app, db)
  NewHandlers(repo)
  helpers.NewHelpers(&app)


  return nil
}

func routes() http.Handler {

	app.Mux.HandleFunc("/", Repo.Home)
	app.Mux.HandleFunc("/user/login", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Repo.GetLogin(w, r)
		case http.MethodPost:
			Repo.PostLogin(w, r)
		}
	}))
	app.Mux.HandleFunc("/error", Repo.ErrorPage)
	app.Mux.HandleFunc("/user/logout", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Repo.GetLogout(w, r)
		case http.MethodPost:
			Repo.PostLogout(w, r)
		}
	}))
	app.Mux.HandleFunc("/user/profile/", Repo.Profile)
	app.Mux.HandleFunc("/post/", Repo.GetPost)
	app.Mux.HandleFunc("/user/post", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			Repo.GetNewPost(w, r)
		case http.MethodPost:
			Repo.PostNewPost(w, r)
		}
	}))
	postPath := "/user/post/"
	app.Mux.Handle(postPath, http.StripPrefix(postPath, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  	switch r.Method {
	  case http.MethodGet:
		  Repo.GetUpdatePost(w, r)
	  case http.MethodPost:
		  Repo.PostUpdatePost(w, r)
	  }
  })))

	return app.Mux

}
