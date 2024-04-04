package repository

import (
	"database/sql"

	"github.com/isadoravieira/api-unisync/src/internal/domain/entity"
)

// Represents an user repository
type User struct {
	db *sql.DB
}

// Create a new user repository
func NewUserRepository(db *sql.DB) *User {
	return &User{db}
}

func (u User) Create(user entity.User) (entity.User, error) {
	return entity.User{}, nil
}