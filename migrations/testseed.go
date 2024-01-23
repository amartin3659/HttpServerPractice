package migrations

import (
	"time"

	"github.com/amartin3659/HttpServerPractice/internal/models"
)

func (m *Seed) TestSeed() {
	users := []models.User{
		{
			ID:        "550e8400-e29b-41d4-a716-446655440010",
			Name:      "Valid User",
			Email:     "valid@text.com",
			Password:  "$2a$10$MT6AiUtj8GcJ11zcIDO/felAezEpqeSseW0OeBa20gGFY/wLKLVG6",
			Role:      models.Admin,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440001",
			Name:      "Invalid User",
			Email:     "invalid@text.com",
			Password:  "$2a$10$MT6AiUtj8GcJ11zcIDO/felAezEpqeSseW0OeBa20gGFY/wLKLVG6",
			Role:      models.Admin,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	m.DB.Users = append(m.DB.Users, users...)
}
