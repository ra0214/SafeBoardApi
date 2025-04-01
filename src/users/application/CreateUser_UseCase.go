package application

import (
	"apiMulti/src/users/domain"

	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	db domain.IUser
}

func NewCreateUser(db domain.IUser) *CreateUser {
	return &CreateUser{db: db}
}

func (cu *CreateUser) Execute(userName string, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	err = cu.db.SaveUser(userName, email, string(hashedPassword))
	if err != nil {
		return err
	}

	return nil
}
