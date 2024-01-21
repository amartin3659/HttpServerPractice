package dbrepo

import (
	"errors"

	"github.com/amartin3659/HttpServerPractice/internal/models"
	"github.com/google/uuid"
)

func (m *mockRepo) Login() {}

func (m *mockRepo) Logout() {}

func (m *mockRepo) GetUserByID(id string) (models.User, error) {
  for _, u := range m.DB.Users {
    if u.ID == id {
      return u, nil
    } 
  }
  return models.User{}, errors.New("No user found")
}

func (m *mockRepo) GetUserByEmail(email string) (models.User, error) {
  for _, u := range m.DB.Users {
    if u.Email == email {
      return u, nil
    }
  }

  return models.User{}, errors.New("User not found")
}

func (m *mockRepo) GetAllPosts() ([]models.Post, error) {
  return m.DB.Posts, nil
}

func (m *mockRepo) GetPostsByUserID(id string) ([]models.Post, error) {
  usersPosts := []models.Post{}
  for _, post := range m.DB.Posts {
    if post.UserID == id {
      usersPosts = append(usersPosts, post)
    }
  }

  return usersPosts, nil
}

func (m *mockRepo) GetPostByID(id string) (models.Post, error) {
  for _, post := range m.DB.Posts {
    pid, err := uuid.Parse(id)
    if err != nil {
      return models.Post{}, errors.New("Cannot parse id")
    }
    if post.ID == pid {
      return post, nil
    }
  }
  return models.Post{}, errors.New("Post not found")
}

func (m *mockRepo) AddPost(models.Post) error {
  return nil
}

func (m *mockRepo) UpdatePost(models.Post) (models.Post, error) {
  return models.Post{}, nil
}
