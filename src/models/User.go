package models

import (
	"api/src/security"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedIn time.Time `json:"createdIn,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if erro := user.validate(stage); erro != nil {
		return erro
	}

	if erro := user.format(stage); erro != nil {
		return erro
	}
	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("field 'Name' is required")
	}

	if user.Nick == "" {
		return errors.New("field 'Nick' is required")
	}

	if user.Email == "" {
		return errors.New("field 'e-mail' is required and cannot be blank")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("the inserted e-mail is invalid")
	}

	if stage == "register" && user.Password == "" {
		return errors.New("field 'Password' is required")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		hashPassword, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}

		user.Password = string(hashPassword)
	}

	return nil
}
