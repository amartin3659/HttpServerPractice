package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/amartin3659/HttpServerPractice/internal/config"
	"github.com/amartin3659/HttpServerPractice/internal/driver"
	"github.com/amartin3659/HttpServerPractice/internal/helpers"
	"github.com/amartin3659/HttpServerPractice/internal/repository"
	"github.com/amartin3659/HttpServerPractice/internal/repository/dbrepo"
)

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

var Repo *Repository

func NewRepo(app *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: app,
		DB:  dbrepo.NewMockRepo(db),
	}
}

func NewHandlers(repo *Repository) {
	Repo = repo
}

type postsResponse struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

// formatAndWrite takes any data interface and a response writer, formats the data into json and writes
func formatAndWrite(input any, w http.ResponseWriter) {
	output, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	posts, err := m.DB.GetAllPosts()
	if err != nil {
	}
  var resp = []postsResponse{}
  for _, post := range posts {
    p := postsResponse{
      Title: post.Title,
      Body: post.Body,
    }
    resp = append(resp, p)
  }
	formatAndWrite(resp, w)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login")
}

func Profile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Profile")
}
