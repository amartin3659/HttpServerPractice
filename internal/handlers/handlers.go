package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	ID     string `json:"id"`
}

// formatAndWrite takes any data interface and a response writer, formats the data into json and writes
func formatAndWrite(input any, w http.ResponseWriter) {
	output, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
	}
	w.Header().Set("Content=Type", "application/json")
	w.Write(output)
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	posts, err := m.DB.GetAllPosts()
	if err != nil {
	}
	var resp = []postsResponse{}
	for _, post := range posts {
		p := postsResponse{
			UserID: post.UserID,
			Title:  post.Title,
			Body:   post.Body,
			ID:     post.ID.String(),
		}
		resp = append(resp, p)
	}
	formatAndWrite(resp, w)
	return
}

func (m *Repository) GetLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login")
}

func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {

}

func (m *Repository) GetLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Logout")
}

func (m *Repository) PostLogout(w http.ResponseWriter, r *http.Request) {

}

type userResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (m *Repository) Profile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	userID := strings.Split(path, "/")[3]
	user, err := m.DB.GetUserByID(userID)
	if err != nil {
		fmt.Println("Error getting user")
	}
	resp := userResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	formatAndWrite(resp, w)
	fmt.Fprint(w, "\n==========================================================================\n")
	posts, err := m.DB.GetPostsByUserID(userID)
	var p = []postsResponse{}
	for _, post := range posts {
		p = append(p, postsResponse{
			UserID: post.UserID,
			Title:  post.Title,
			Body:   post.Body,
			ID:     post.ID.String(),
		})
	}
	formatAndWrite(p, w)
	return
}

func (m *Repository) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	postID := strings.Split(path, "/")[2]
	post, err := m.DB.GetPostByID(postID)
	if err != nil {
		helpers.ServerError(w, err)
		fmt.Fprint(w, "There was an error")
	}
	resp := postsResponse{
		UserID: post.UserID,
		Title:  post.Title,
		Body:   post.Body,
		ID:     post.ID.String(),
	}
	formatAndWrite(resp, w)
}

func (m *Repository) PostPost(w http.ResponseWriter, r *http.Request) {

}

func (m *Repository) UpdatePost(w http.ResponseWriter, r *http.Request) {

}
