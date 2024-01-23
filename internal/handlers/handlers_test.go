package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/amartin3659/HttpServerPractice/internal/driver"
	"github.com/amartin3659/HttpServerPractice/internal/models"
	"github.com/google/uuid"
)

func TestNewRepo(t *testing.T) {
  var testdb driver.DB
  testRepo := NewRepo(&app, &testdb)

  if reflect.TypeOf(testRepo).String() != "*handlers.Repository" {
    t.Error("Types did not match")
  }
}

func TestHome(t *testing.T) {
  // No post
  req := httptest.NewRequest("GET", "/home", nil)
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(Repo.Home)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusTemporaryRedirect {
    t.Error("Unexpected code", rr.Code)
  }

  // OK
  // add a post
  nPost := models.Post{
    ID: uuid.New(),
    Title: "Test Post",
    Body: "This is a test post",
    UserID: "550e8400-e29b-41d4-a716-446655440010",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  Repo.DB.AddPost(nPost)
  req = httptest.NewRequest("GET", "/home", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.Home)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusOK {
    t.Error("Expected different code")
  }
}

func TestLogin(t *testing.T) {
  // GetLogin - OK
  req := httptest.NewRequest("GET", "/user/login", nil)
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(Repo.GetLogin)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusOK {
    t.Error("Expected a different code")
  }

  // PostLogin - OK
  postData := url.Values{}
  postData.Add("email", "valid@text.com")
  postData.Add("password", "password")
  req = httptest.NewRequest("POST", "/user/login", strings.NewReader(postData.Encode()))
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.PostLogin)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusSeeOther {
    t.Error("Expected a different code")
  }
  // check if cookie was set
  sessionCookie := "session"
  var targetCookie *http.Cookie
  cookies := rr.Result().Cookies()
  for _, cookie := range cookies {
    if cookie.Name == sessionCookie {
      targetCookie = cookie
      break
    }
  }
  if targetCookie == nil {
    t.Error("Session cookie was not set")
  } 

  // Now that we have a cookie test edge cases
  // GetLogin - With valid token
  req = httptest.NewRequest("GET", "/user/login", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.GetLogin)
  req.AddCookie(targetCookie)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusSeeOther {
    t.Error("Expected a different code")
  }

  // PostLogin - Invalid form
  // has to be http.NewRequest to test form error
  req, _ = http.NewRequest("POST", "/user/login", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.PostLogin)
//  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusUnauthorized {
    t.Error("Expected a different code")
  }

  // PostLogin - Empty form
  req = httptest.NewRequest("POST", "/user/login", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.PostLogin)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusUnauthorized {
    t.Error("Expected a different code")
  }

  // PostLogin - User doesn't exist
  postData = url.Values{}
  postData.Add("email", "nouser@text.com")
  postData.Add("password", "password")
  req = httptest.NewRequest("POST", "/user/login", strings.NewReader(postData.Encode()))
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.PostLogin)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusUnauthorized {
    t.Error("Expected a different code")
  }

  // PostLogin - Incorrect password
  postData = url.Values{}
  postData.Add("email", "valid@text.com")
  postData.Add("password", "pssword")
  req = httptest.NewRequest("POST", "/user/login", strings.NewReader(postData.Encode()))
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.PostLogin)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusUnauthorized {
    t.Error("Expected a different code")
  }
}

func TestErrorPage(t *testing.T) {
  req := httptest.NewRequest("GET", "/error", nil)
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(Repo.ErrorPage)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusOK {
    t.Error("Was expecting a different code")
  }
}

func TestGetLogout(t *testing.T) {
  req := httptest.NewRequest("GET", "/user/logout", nil)
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(Repo.GetLogout)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusOK {
    t.Error("Expected a different code")
  }
}

func TestPostLogout(t *testing.T) {
  // PostLogout - OK
  // Login first
  postData := url.Values{}
  postData.Add("email", "valid@text.com")
  postData.Add("password", "password")
  req := httptest.NewRequest("POST", "/user/login", strings.NewReader(postData.Encode()))
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(Repo.PostLogin)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusSeeOther {
    t.Error("Expected a different code")
  }
  var sessionCookie *http.Cookie
  for _, cookie := range rr.Result().Cookies() {
    if cookie.Name == "session" {
      sessionCookie = cookie
      break
    }
  }
  if sessionCookie == nil {
    t.Error("Session cookie was not set")
  }
  // now test logout
  req = httptest.NewRequest("POST", "/user/logout", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.PostLogout)
  req.AddCookie(sessionCookie)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusSeeOther {
    t.Error("Expecting a different code")
  }

  // PostLogout - No session
  req = httptest.NewRequest("POST", "/user/logout", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.PostLogout)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusSeeOther {
    t.Error("Expecting a different code")
  }
}

func TestProfile(t *testing.T) {
  // Profile - OK
  // login first
  postData := url.Values{}
  postData.Add("email", "valid@text.com")
  postData.Add("password", "password")
  req := httptest.NewRequest("POST", "/user/login", strings.NewReader(postData.Encode()))
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(Repo.PostLogin)
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusSeeOther {
    t.Error("Expected a different code")
  }
  var sessionCookie *http.Cookie
  for _, cookie := range rr.Result().Cookies() {
    if cookie.Name == "session" {
      sessionCookie = cookie
      break
    }
  }
  if sessionCookie == nil {
    t.Error("Session cookie was not set")
  }
  req = httptest.NewRequest("GET", "/user/profile/550e8400-e29b-41d4-a716-446655440010", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.Profile)
  req.AddCookie(sessionCookie)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusOK {
    t.Error("Expected a different code")
  }

  // Profile - User not found
  req = httptest.NewRequest("GET", "/user/profile/450e8400-e29b-41d4-a716-446655440010", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.Profile)
  req.AddCookie(sessionCookie)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusTemporaryRedirect {
    t.Error("Expected a different code")
  }
}

func TestGetPost(t *testing.T) {
  pid := uuid.New()
  nPost := models.Post{
    ID: pid,
    Title: "Test Post",
    Body: "This is a test post",
    UserID: "550e8400-e29b-41d4-a716-446655440010",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }
  Repo.DB.AddPost(nPost)
  // GetPost - OK
  req := httptest.NewRequest("GET", "/post/"+pid.String(), nil)
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(Repo.GetPost)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusOK {
    t.Error("Expected a different code")
  }

  // GetPost - No post
  req = httptest.NewRequest("GET", "/post/550e8401-e29b-41d4-a716-443655440060", nil)
  rr = httptest.NewRecorder()
  handler = http.HandlerFunc(Repo.GetPost)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusTemporaryRedirect {
    t.Error("Expected a different code")
  }
}

