package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
  os.Exit(m.Run())
}

func TestRun(t *testing.T) {
  err := run()
  if err != nil {
    t.Error("Main encountered an error")
  }
}
