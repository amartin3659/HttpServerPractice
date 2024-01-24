package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/amartin3659/HttpServerPractice/internal/config"
	"github.com/amartin3659/HttpServerPractice/internal/driver"
	"github.com/amartin3659/HttpServerPractice/internal/helpers"
	"github.com/amartin3659/HttpServerPractice/internal/models"
	"github.com/amartin3659/HttpServerPractice/internal/repository"
	"github.com/amartin3659/HttpServerPractice/internal/repository/dbrepo"
	"github.com/amartin3659/HttpServerPractice/internal/session"
	"github.com/amartin3659/HttpServerPractice/internal/templates"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func NewTestRepo(app *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: app,
		DB:  dbrepo.NewTestRepo(db),
	}
}

func NewHandlers(repo *Repository) {
	Repo = repo
}

// formatAndWrite takes any data interface and a response writer, formats the data into json and writes
func formatAndWrite(input any, w http.ResponseWriter) {
	output, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		helpers.ServerError(w, err)
    return
	}
	w.Header().Set("Content=Type", "application/json")
	w.Write(output)
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	posts, err := m.DB.GetAllPosts()
	if err != nil {
    http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
    return
	}
	var resp = []models.PostsResponse{}
	for _, post := range posts {
		p := models.PostsResponse{
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
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("No session. Login")
	}
	if cookie.Valid() == nil {
		userID, err := m.App.Session.Get(uuid.MustParse(cookie.Value))
		if err == nil {
			http.Redirect(w, r, "/user/profile/"+userID.String(), http.StatusSeeOther)
		}
    fmt.Println("Session not found")
	}
	htmlTemplate := templates.LoginTemplate()
	tmpl, err := template.New("Login").Parse(htmlTemplate)
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (m *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form data")
		http.Redirect(w, r, "/error", http.StatusUnauthorized)
		return
	}
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	if email == "" || password == "" {
		fmt.Println("No email or password")
		http.Redirect(w, r, "/error", http.StatusUnauthorized)
		return
	}
	// get hashed password
	user, err := m.DB.GetUserByEmail(email)
	if err != nil {
		fmt.Println("No User")
		http.Redirect(w, r, "/error", http.StatusUnauthorized)
		return
	}
	hashedPass := user.Password
	// check if correct password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("Incorrect Password")
		http.Redirect(w, r, "/error", http.StatusUnauthorized)
		return
	}
	// create token and session etc...
	sessionID, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("Could not set UUID")
		http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
	}
	userID := uuid.MustParse(user.ID)
	session := session.Session{
		SessionID: sessionID,
		UserID:    userID,
	}
	m.App.Session.Add(session)
	cookie := http.Cookie{
		Name:    "session",
		Value:   sessionID.String(),
		MaxAge:  120,
		Path:    "/",
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/user/profile/"+user.ID, http.StatusSeeOther)
	return
}

func (m *Repository) ErrorPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "There was an error authenticating your credentials")
}

func (m *Repository) GetLogout(w http.ResponseWriter, r *http.Request) {
	htmlTemplate := templates.GetLogout()
	tmpl, err := template.New("Logout").Parse(htmlTemplate)
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (m *Repository) PostLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("No session data")
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
    return
	}
  sessionID := uuid.MustParse(cookie.Value)
	expireCookie := http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
    Path: "/",
	}
	http.SetCookie(w, &expireCookie)
	m.App.Session.Remove(sessionID)
	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

func (m *Repository) Profile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	userID := strings.Split(path, "/")[3]
	user, err := m.DB.GetUserByID(userID)
	if err != nil {
		fmt.Println("Error getting user")
    http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
    return
	}
	resp := models.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	formatAndWrite(resp, w)
	fmt.Fprint(w, "\n==========================================================================\n")
	posts, err := m.DB.GetPostsByUserID(userID)
	var p = []models.PostsResponse{}
	for _, post := range posts {
		p = append(p, models.PostsResponse{
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
    http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
    return
	}
	resp := models.PostsResponse{
		UserID: post.UserID,
		Title:  post.Title,
		Body:   post.Body,
		ID:     post.ID.String(),
	}
	formatAndWrite(resp, w)
  return
}

func (m *Repository) GetNewPost(w http.ResponseWriter, r *http.Request) {
  // must be logged into make post
  cookie, err := r.Cookie("session")
  if err != nil {
    fmt.Println("Need to be logged in to make post")
    // redirect to login
    http.Redirect(w, r, "/user/login", http.StatusTemporaryRedirect)
    return
  }
  userID, _ := m.App.Session.Get(uuid.MustParse(cookie.Value))
  if uuid.Nil == userID {
    fmt.Println("Session is not valid")
    http.Redirect(w, r, "/user/login", http.StatusTemporaryRedirect)
    return
  }
	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
    <title>New Post</title>
</head>
<body>
  <form method="POST" action="/user/post" novalidate>
      <label for="title">Title:</lable>
      <br>
      <input type="text" name="title" />
      <br>
      <br>
      <label for="body">Body:</label>
      <br>
      <textarea name="body" rows="10" cols="30"></textarea>
      <br>
      <br>
      <button type="submit">Create</button>
      <br>
      <br>
    </form>
</body>
</html>
`
	tmpl, err := template.New("NewPost").Parse(htmlTemplate)
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (m *Repository) PostNewPost(w http.ResponseWriter, r *http.Request) {
  // add new post with userID
  // if session ended between creating post and submitting it, just redirect to login page
  cookie, err := r.Cookie("session")
  if err != nil {
    fmt.Println("Need to be logged in to make post")
    // redirect to login
    http.Redirect(w, r, "/user/login", http.StatusTemporaryRedirect)
    return
  }
  userID, _ := m.App.Session.Get(uuid.MustParse(cookie.Value))
  if uuid.Nil == userID {
    fmt.Println("Session is not valid")
    http.Redirect(w, r, "/user/login", http.StatusTemporaryRedirect)
    return
  }
	err = r.ParseForm()
	if err != nil {
		fmt.Println("Error parsing form data")
		http.Redirect(w, r, "/error", http.StatusUnauthorized)
		return
	}
  title := r.PostForm.Get("title")
  body := r.PostForm.Get("body")
  newPost := models.Post{
    ID: uuid.New(),
    UserID: userID.String(),
    Title: title,
    Body: body,
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  m.DB.AddPost(newPost)
  http.Redirect(w, r, "/post/"+newPost.ID.String(), http.StatusSeeOther)
  return
}

func (m *Repository) GetUpdatePost(w http.ResponseWriter, r *http.Request) {
  // again need to be logged in, redirect if not logged in
  // grab post id from url, if doesn't exist display error
  // if exists populate fields with post data
  postID := strings.Split(r.URL.Path, "/")[3]
	post, err := m.DB.GetPostByID(postID)
	if err != nil {
		fmt.Println("Error getting post")
    http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
    return
	}
  userID := uuid.MustParse(post.UserID)
  cookie, err := r.Cookie("session")
  if err != nil {
    fmt.Println("Need to be logged in to make post")
    // redirect to login
    http.Redirect(w, r, "/user/login", http.StatusTemporaryRedirect)
    return
  }
  cUserID, _ := m.App.Session.Get(uuid.MustParse(cookie.Value))
  if uuid.Nil == cUserID {
    fmt.Println("Session is not valid")
    http.Redirect(w, r, "/user/login", http.StatusTemporaryRedirect)
    return
  }
  if userID != cUserID {
    fmt.Println("Can only edit your own posts")
    http.Redirect(w, r, "/home", http.StatusSeeOther)
    return
  }
	htmlTemplate := templates.PostTemplate(postID)
	tmpl, err := template.New("UpdatePost").Parse(htmlTemplate)
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (m *Repository) PostUpdatePost(w http.ResponseWriter, r *http.Request) {
  // redirect to login page if session is expired
  // update post with new contents
  path := r.URL.Path
  postID := strings.Split(path, "/")[3]
	post, err := m.DB.GetPostByID(postID)
	if err != nil {
		fmt.Println("Error getting post")
    http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
    return
	}
  userID := uuid.MustParse(post.UserID)
  cookie, err := r.Cookie("session")
  if err != nil {
    fmt.Println("Need to be logged in to make post")
    // redirect to login
    http.Redirect(w, r, "/user/login", http.StatusTemporaryRedirect)
    return
  }
  cUserID, _ := m.App.Session.Get(uuid.MustParse(cookie.Value))
  if uuid.Nil == cUserID {
    fmt.Println("Session is not valid")
    http.Redirect(w, r, "/user/login", http.StatusTemporaryRedirect)
    return
  }
  if userID != cUserID {
    fmt.Println("Can only edit your own posts")
    http.Redirect(w, r, "/home", http.StatusSeeOther)
    return
  }
  err = r.ParseForm()
  if err != nil {
    fmt.Println("Could not parse form")
    http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
    return
  }
  title := r.PostForm.Get("title")
  body := r.PostForm.Get("body")
  nPost := models.Post{
    ID: uuid.MustParse(postID),
    UserID: userID.String(),
    Title: title,
    Body: body,
    CreatedAt: post.CreatedAt,
    UpdatedAt: time.Now(),
  }
  uPost, err := m.DB.UpdatePost(nPost)
  if err != nil {
    fmt.Println("Could not update")
    http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
    return
  }
  http.Redirect(w, r, "/post/"+uPost.ID.String(), http.StatusSeeOther)
  return
}
