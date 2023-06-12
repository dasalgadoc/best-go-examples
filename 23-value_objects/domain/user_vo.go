package domain

import "time"

type UserVo struct {
	id        UserId
	email     UserEmail
	birthDate UserBirthDate
}

// Constructor with primitive types, Value objects and modeling are not exposed
func NewUserVo(id, email string, birthDate time.Time) (*UserVo, error) {
	userId, err := NewUserId(id)
	if err != nil {
		return nil, err
	}
	userEmail, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}
	userBirthDate, err := NewUserBirthDate(birthDate)
	if err != nil {
		return nil, err
	}

	return &UserVo{
		id:        *userId,
		email:     *userEmail,
		birthDate: *userBirthDate,
	}, nil
}

func (u *UserVo) UpdateEmail(newEmail string) error {
	userEmail, err := NewUserEmail(newEmail)
	if err != nil {
		return err
	}
	u.email = *userEmail
	return nil
}

/*--Getters--*/
func (u *UserVo) Id() string {
	return u.id.Value()
}

func (u *UserVo) Email() string {
	return u.email.Value()
}

func (u *UserVo) BirthDate() time.Time {
	return u.birthDate.Value()
}
