package driver

import (
	"github.com/amartin3659/HttpServerPractice/internal/models"
)

type DB struct {
	Users []models.User
	Posts []models.Post
}

func NewDB() *DB {
	return &DB{}
}
