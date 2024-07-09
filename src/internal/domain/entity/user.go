package entity

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/isadoravieira/api-unisync/src/pkg/security"
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
func (user *User) PrepareUser(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.formatFields(stage); err != nil {
		return err 
	}

	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("Name is a required field and cannot be blank")
	}

	if user.UserName == "" {
		return errors.New("UserName is a required field and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("Email is a required field and cannot be blank")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email entered is an invalid format")
	}

	if stage == "register" && user.Password == "" {
		return errors.New("Password is a required field and cannot be blank")
	}

	return nil
}

func (user *User) formatFields(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.UserName = strings.TrimSpace(user.UserName)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		passwordHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordHash)
	}
	return nil
}
