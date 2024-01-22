package dbrepo

import (
	"errors"
	"fmt"

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
      fmt.Println("Cannot parse uuid")
      return models.Post{}, errors.New("Cannot parse id")
    }
    if post.ID == pid {
      return post, nil
    }
  }
  return models.Post{}, errors.New("Post not found")
}

func (m *mockRepo) AddPost(post models.Post) error {
  m.DB.Posts = append(m.DB.Posts, post)
  return nil
}

func (m *mockRepo) UpdatePost(p models.Post) (models.Post, error) {
  for i, post := range m.DB.Posts {
    if post.ID == p.ID {
      m.DB.Posts[i] = p 
      return p, nil
    }
  }
  return models.Post{}, errors.New("Could not update")
}
