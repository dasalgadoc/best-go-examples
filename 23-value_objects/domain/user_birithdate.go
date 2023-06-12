package domain

import (
	"errors"
	"time"
)

type UserBirthDate struct {
	value time.Time
}

func NewUserBirthDate(value time.Time) (*UserBirthDate, error) {
	err := ensureBirthDateIsValidVO(value)
	if err != nil {
		return nil, err
	}
	return &UserBirthDate{value: value}, nil
}

func ensureBirthDateIsValidVO(birthDate time.Time) error {
	// ... some validations
	if birthDate.IsZero() {
		return errors.New("birth date is required")
	}
	currentDate := time.Now()
	if birthDate.After(currentDate) {
		return errors.New("birth date is invalid")
	}
	currentYear := currentDate.Year()
	year := birthDate.Year()
	if currentYear-year < 18 {
		return errors.New("user must be at least 18 years old")
	}
	if currentYear-year > 100 {
		return errors.New("user must be at most 100 years old")
	}

	return nil
}

func (u *UserBirthDate) Value() time.Time {
	return u.value
}
