package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestRoutes(t *testing.T) {
  mux := routes()

  switch v := mux.(type) {
  case http.Handler:
    // all fine nothing to do
  default:
    t.Error(fmt.Sprintf("Type mismatch: Expected http.Handler, got %T", v))
  }
}

