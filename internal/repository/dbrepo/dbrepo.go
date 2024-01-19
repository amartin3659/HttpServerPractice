package dbrepo

import (
	"github.com/amartin3659/HttpServerPractice/internal/driver"
	"github.com/amartin3659/HttpServerPractice/internal/repository"
)

type mockDBRepo struct {
  DB *driver.DB
}

func NewMockRepo(conn *driver.DB) repository.DatabaseRepo {
  return &mockDBRepo{
    DB: conn,
  }
}
