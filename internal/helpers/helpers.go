package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/amartin3659/HttpServerPractice/internal/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
  app = a
}

func ServerInfo(message string) {
  app.InfoLog.Println(message)
}

func ServerError(w http.ResponseWriter, err error) {
  trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
  app.ErrorLog.Println(trace)
  http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError) 
}
