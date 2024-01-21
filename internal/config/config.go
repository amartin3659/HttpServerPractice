package config

import (
	"log"
	"net/http"
	"os"

	"github.com/amartin3659/HttpServerPractice/internal/session"
)

type AppConfig struct {
	InProduction bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	Mux          *http.ServeMux
	Session      *session.Sessions
}

func init() {
	os.Setenv("port", ":8080")
}

func (app *AppConfig) SetInProduction(inProduction bool) {
	app.InProduction = inProduction
}

func (app *AppConfig) SetInfoLog() {
	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog
}

func (app *AppConfig) SetErrorLog() {
	errorLog := log.New(os.Stdout, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.InfoLog = errorLog
}

func (app *AppConfig) SetMux(mux *http.ServeMux) {
	app.Mux = mux
}

func (app *AppConfig) SetSession(session *session.Sessions) {
  app.Session = session
}
