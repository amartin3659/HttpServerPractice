package repository

import (
	"github.com/amartin3659/HttpServerPractice/internal/models"
)

// Setup similar to VacationHomes project
// Use a mock db (slices) to fake data
// Have option then to use real db in future
type DatabaseRepo interface {
  Login()
  Logout()
  GetUserByID(id string) (models.User, error)
  GetAllPosts() ([]models.Post, error)
  GetPostsByUserID(id string) ([]models.Post, error)
  GetPostByID(id string) (models.Post, error)
  AddPost(models.Post) (error)
  UpdatePost(models.Post) (models.Post, error)
}
