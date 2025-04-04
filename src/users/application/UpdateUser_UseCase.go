package application

import (
	"apiMulti/src/users/domain"
)

type EditUser struct {
	db domain.IUser
}

func NewEditUser(db domain.IUser) *EditUser {
	return &EditUser{db: db}
}

func (eu *EditUser) Execute(id int32, userName string, email string, password string, esp32ID string) error {
	return eu.db.UpdateUser(id, userName, email, password, esp32ID)
}