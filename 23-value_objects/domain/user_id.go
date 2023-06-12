package domain

import "errors"

type UserId struct {
	value string
}

func NewUserId(value string) (*UserId, error) {
	err := ensureIdIsValidVO(value)
	if err != nil {
		return nil, err
	}
	return &UserId{value: value}, nil
}

func ensureIdIsValidVO(id string) error {
	// ... some validations
	if id == "" {
		return errors.New("id is required")
	}
	return nil
}

func (u *UserId) Value() string {
	return u.value
}
