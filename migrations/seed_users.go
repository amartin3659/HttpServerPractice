package migrations

import (
	"time"

	"github.com/amartin3659/HttpServerPractice/internal/models"
)

func (m *Seed) seedUsers() {
	users := []models.User{
		{
			ID:        "550e8400-e29b-41d4-a716-446655440010",
			Name:      "Alex Martin",
			Email:     "amartin.code@gmail.com",
			Password:  "$2a$10$WXdfUSKy0shI/yJ15CUgh.Tt72tNYwFGlNBHxark5uVlZNbyzeyga",
			Role:      models.Admin,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440000",
			Name:      "Emily Johnson",
			Email:     "emily7892@test.com",
			Password:  "$2a$10$kUB5bmH75UonMwHLIJOpH.xBKkwvJvSIKRHwcjtpYO4kPZbFYNM0i",
			Role:      models.Regular,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440001",
			Name:      "Alex Rodriguez",
			Email:     "alex2156@test.com",
			Password:  "$2a$10$81OP.jOl6Jb6Pvny2ELmjePP.f0kluMTng84AwS7N8FGpcA5AlBWq",
			Role:      models.Regular,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440002",
			Name:      "Sarah Patel",
			Email:     "sarah3764@test.com",
			Password:  "$2a$10$nYTTBaiuU6nH.MarkYvUE.17sB5XkJ/z5ZmU7gJcQSaPJQ1mKXPx6",
			Role:      models.Regular,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440003",
			Name:      "Kevin Chen",
			Email:     "kevin9021@test.com",
			Password:  "$2a$10$bbJKvfSry6pHSpoXgnTdeei/Q/FNBBYXPmZHHEtjmngv8VFatKgZK",
			Role:      models.Regular,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440004",
			Name:      "Jasmine Carter",
			Email:     "jasmine1348@test.com",
			Password:  "$2a$10$qaqqNN3k0xe4nCm1eJJ5LeZsZsa/.1EOpJWAWyjDzb8V5FpiSoy6G",
			Role:      models.Regular,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440005",
			Name:      "Ryan Thompson",
			Email:     "ryan5673@test.com",
			Password:  "$2a$10$U4g3OjrZ5kN9J1bNFo6AWOKkJbMC0o6RzlV8oC.HekWGI6cbfLbrC",
			Role:      models.Regular,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440006",
			Name:      "Olivia Davis",
			Email:     "olivia9987@test.com",
			Password:  "$2a$10$dkDgMNmWVRjeJRpIlUU03eB9j/fFM5YQ8rcGqI515AxP5onfCEbJe",
			Role:      models.Regular,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440007",
			Name:      "Marcus White",
			Email:     "marcus4523@test.com",
			Password:  "$2a$10$w5dDuitesPdpqSsf55DKROVris/XpF7Ki5jLN8MPEtwQFZTm84asu",
			Role:      models.Regular,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440008",
			Name:      "Lauren Miller",
			Email:     "lauren7890@test.com",
			Password:  "$2a$10$Ql5Z3ffPh.D9HHOVep570u0RVMnbv1K6o5lr7.t/DMAOY0bCqSope",
			Role:      models.Regular,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "550e8400-e29b-41d4-a716-446655440009",
			Name:      "Brandon Lee",
			Email:     "brandon3210@test.com",
			Password:  "$2a$10$LWPnKSa.xwSjlz0iJSonguU/RVzOf4PxBeXY/c/M8KMkWP7z9BMfO",
			Role:      models.Regular,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	m.DB.Users = append(m.DB.Users, users...)
}
