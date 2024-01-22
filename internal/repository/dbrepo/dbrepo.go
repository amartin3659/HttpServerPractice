package dbrepo

import (
	"github.com/amartin3659/HttpServerPractice/internal/driver"
	"github.com/amartin3659/HttpServerPractice/internal/repository"
)

type mockRepo struct {
  DB *driver.DB
}

// this part confuses me DB is a pointer but returning an interface
func NewMockRepo(conn *driver.DB) repository.DatabaseRepo {
  return &mockRepo{
    DB: conn,
  }
}
