package dbrepo

import (
	"errors"

	"github.com/amartin3659/HttpServerPractice/internal/models"
)

func (m *mockDBRepo) GetUserByID(id string) (models.User, error) {
  for _, u := range m.DB.Users {
    if u.ID == id {
      return u, nil
    } 
  }
  return models.User{}, errors.New("No user found")
}

func (m *mockDBRepo) GetAllPosts() ([]models.Post, error) {
  // link user to post
  return m.DB.Posts, nil
}
