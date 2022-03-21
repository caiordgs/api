package models

import (
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

	user.format()
	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("Field 'Name' is required.")
	}

	if user.Nick == "" {
		return errors.New("Field 'Nick' is required.")
	}

	if user.Email == "" {
		return errors.New("Field 'e-mail' is required and cannot be blank.")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("The inserted e-mail is invalid.")
	}

	if stage == "register" && user.Password == "" {
		return errors.New("Field 'Password' is required.")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
