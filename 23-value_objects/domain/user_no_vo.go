package domain

import (
	"errors"
	"regexp"
	"time"
)

// UserNoVO is a struct that represents a user without value objects
type UserNoVO struct {
	id        string
	email     string
	birthDate time.Time
}

func NewUserNoVO(id, email string, birthDate time.Time) (*UserNoVO, error) {
	err := ensureIdIsValid(id)
	if err != nil {
		return nil, err
	}
	err = ensureEmailIsValid(email)
	if err != nil {
		return nil, err
	}
	err = ensureBirthDateIsValid(birthDate)
	if err != nil {
		return nil, err
	}

	return &UserNoVO{
		id:        id,
		email:     email,
		birthDate: birthDate,
	}, nil
}

func (u *UserNoVO) UpdateEmail(newEmail string) error {
	err := ensureEmailIsValid(newEmail)
	if err != nil {
		return err
	}
	u.email = newEmail
	return nil
}

func ensureIdIsValid(id string) error {
	// ... some validations
	if id == "" {
		return errors.New("id is required")
	}
	return nil
}

func ensureEmailIsValid(email string) error {
	// ... some validations
	regularExpression := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	if email == "" {
		return errors.New("email is required")
	}

	if match, _ := regexp.MatchString(regularExpression, email); !match {
		return errors.New("email is invalid")
	}
	return nil
}

func ensureBirthDateIsValid(date time.Time) error {
	// ... some validations
	if date.IsZero() {
		return errors.New("birth date is required")
	}

	currentDate := time.Now()
	if date.After(currentDate) {
		return errors.New("birth date is invalid")
	}

	currentYear := currentDate.Year()
	year := date.Year()
	if currentYear-year < 18 {
		return errors.New("user must be at least 18 years old")
	}
	if currentYear-year > 100 {
		return errors.New("user must be at most 100 years old")
	}

	return nil
}
