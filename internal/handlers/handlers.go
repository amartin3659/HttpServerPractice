package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Home")
}

func Login(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Login")
}

func Profile(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Profile")
}
