package main

import (
	"net/http"
	"os"

	"github.com/amartin3659/HttpServerPractice/internal/config"
	"github.com/amartin3659/HttpServerPractice/internal/driver"
	"github.com/amartin3659/HttpServerPractice/internal/handlers"
	"github.com/amartin3659/HttpServerPractice/internal/helpers"
	"github.com/amartin3659/HttpServerPractice/migrations"
)

var app config.AppConfig

func main() {
  app.SetInProduction(false)
  app.SetInfoLog()
  app.SetErrorLog()
  app.SetMux(http.NewServeMux())

  db := driver.NewDB()
  seed := migrations.NewSeed(db)
  seed.Seed()
  repo := handlers.NewRepo(&app, db)
  handlers.NewHandlers(repo)
  helpers.NewHelpers(&app)

  server := &http.Server{
    Addr: os.Getenv("port"),
    Handler: routes(),
  }

  err := server.ListenAndServe()
  if err != nil {
    app.ErrorLog.Println(err)
  }

}
