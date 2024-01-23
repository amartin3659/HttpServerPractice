package dbrepo

import (
	"errors"
	"fmt"

	"github.com/amartin3659/HttpServerPractice/internal/models"
	"github.com/google/uuid"
)

func (m *testRepo) Login() {}

func (m *testRepo) Logout() {}

func (m *testRepo) GetUserByID(id string) (models.User, error) {
  for _, u := range m.DB.Users {
    if u.ID == id {
      return u, nil
    } 
  }
  return models.User{}, errors.New("No user found")
}

func (m *testRepo) GetUserByEmail(email string) (models.User, error) {
  for _, u := range m.DB.Users {
    if u.Email == email {
      return u, nil
    }
  }

  return models.User{}, errors.New("User not found")
}

func (m *testRepo) GetAllPosts() ([]models.Post, error) {
  if len(m.DB.Posts) == 0 {
    return m.DB.Posts, errors.New("No posts")
  }
  return m.DB.Posts, nil
}

func (m *testRepo) GetPostsByUserID(id string) ([]models.Post, error) {
  usersPosts := []models.Post{}
  for _, post := range m.DB.Posts {
    if post.UserID == id {
      usersPosts = append(usersPosts, post)
    }
  }

  return usersPosts, nil
}

func (m *testRepo) GetPostByID(id string) (models.Post, error) {
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

func (m *testRepo) AddPost(post models.Post) error {
  m.DB.Posts = append(m.DB.Posts, post)
  return nil
}

func (m *testRepo) UpdatePost(p models.Post) (models.Post, error) {
  for i, post := range m.DB.Posts {
    if post.ID == p.ID {
      m.DB.Posts[i] = p 
      return p, nil
    }
  }
  return models.Post{}, errors.New("Could not update")
}
