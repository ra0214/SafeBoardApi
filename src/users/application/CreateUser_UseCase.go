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

func (cu *CreateUser) Execute(userName string, email string, password string, esp32ID string) error {
	// Generar hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Guardar usuario con el ESP32ID
	err = cu.db.SaveUser(userName, email, string(hashedPassword), esp32ID)
	if err != nil {
		return err
	}

	return nil
}
