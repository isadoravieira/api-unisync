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
	statement, err := u.db.Prepare(
		"INSERT INTO users (name, username, email, password) VALUES(?, ?, ?, ?)",
	)
	if err != nil {
		return entity.User{}, err
	}
	defer statement.Close()

	newUser, err := statement.Exec(user.Name, user.UserName, user.Email, user.Password)
	if err != nil {
		return entity.User{}, err
	}

	lastUserID, err := newUser.LastInsertId()
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		ID:        lastUserID,
		Name:      user.Name,
		UserName:  user.UserName,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}, nil
}
