package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
  req := httptest.NewRequest("GET", "/", nil)
  rr := httptest.NewRecorder()
  handler := http.HandlerFunc(Repo.Home)
  handler.ServeHTTP(rr, req)
  if rr.Code != http.StatusOK {
    t.Error("Expected different code")
  }
}
