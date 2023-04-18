package models

import (
	"errors"
	"regexp"
)

type User struct {
	Base
	Username string `json:"username" gorm:"unique;notNull"`
	Password string `json:"password" gorm:"unique;notNull"`
}

func (user User) ValidatePassword() error {
	if user.Password == "" {
		return errors.New("Password must not be empty")
	}

	if len(user.Password) < 8 {
		return errors.New("Password must have at least 8 characters")
	}

	hasNumber, err := regexp.Compile("[0-9]")
	if err != nil {
		return err
	}
	if !hasNumber.MatchString(user.Password) {
		return errors.New("Password must contain at least one number")
	}

	hasLetter, err := regexp.Compile("[a-zA-Z]")
	if err != nil {
		return err
	}
	if !hasLetter.MatchString(user.Password) {
		return errors.New("Password must contain at least one letter")
	}

	return nil
}
