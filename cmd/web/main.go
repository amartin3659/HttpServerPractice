package main

import (
	"net/http"
	"os"

	"github.com/amartin3659/HttpServerPractice/internal/config"
	"github.com/amartin3659/HttpServerPractice/internal/helpers"
)

var app config.AppConfig

func main() {

  app.SetInProduction(false)
  app.SetInfoLog()
  app.SetErrorLog()
  app.SetMux(http.NewServeMux())
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
