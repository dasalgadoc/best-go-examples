package domain

import (
	"errors"
	"regexp"
)

type UserEmail struct {
	value string
}

func NewUserEmail(value string) (*UserEmail, error) {
	err := emailIsNotEmpty(value)
	if err != nil {
		return nil, err
	}
	err = emailValid(value)
	if err != nil {
		return nil, err
	}
	return &UserEmail{value: value}, nil
}

func (u *UserEmail) Value() string {
	return u.value
}

func emailValid(email string) error {
	regularExpression := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	if match, _ := regexp.MatchString(regularExpression, email); !match {
		return errors.New("email is invalid")
	}
	return nil
}

func emailIsNotEmpty(email string) error {
	if email == "" {
		return errors.New("email is required")
	}
	return nil
}
