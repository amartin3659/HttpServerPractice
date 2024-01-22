package main

import "testing"

func TestAuth(t *testing.T) {
  ok := auth()
  if !ok {
    t.Error("Error") 
  }
}
