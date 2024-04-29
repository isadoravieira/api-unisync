package entity

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	UserName  string    `json:"userName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare will call validate and format methods for user received
func (user *User) PrepareUser() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.formatFields()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("Name is a required field and cannot be blank")
	}

	if user.UserName == "" {
		return errors.New("UserName is a required field and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("Email is a required field and cannot be blank")
	}

	if user.Password == "" {
		return errors.New("Password is a required field and cannot be blank")
	}

	return nil
}

func (user *User) formatFields() {
	user.Name = strings.TrimSpace(user.Name)
	user.UserName = strings.TrimSpace(user.UserName)
	user.Email = strings.TrimSpace(user.Email)
}
