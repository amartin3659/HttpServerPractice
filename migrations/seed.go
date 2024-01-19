package migrations

import "github.com/amartin3659/HttpServerPractice/internal/driver"

type Seed struct {
  DB *driver.DB
}

func NewSeed(db *driver.DB) Seed {
  return Seed{
    DB: db,
  }
}

func (m *Seed) Seed() {
  m.seedUsers()
  m.seedPosts()
}
