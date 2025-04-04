package application

import (
	"apiMulti/src/movement/domain"
)

type CreateMovement struct {
	rabbit domain.IMovementRabbitqm
	db     domain.IMovement
}

func NewCreateMovement(r domain.IMovementRabbitqm, db domain.IMovement) *CreateMovement {
	return &CreateMovement{rabbit: r, db: db}
}

func (ct *CreateMovement) Execute(esp32_id string, aceleracion float64) error {
	err := ct.db.SaveMovement(esp32_id, aceleracion)
	if err != nil {
		return err
	}

	movement := domain.NewMovement(esp32_id, aceleracion)

	err = ct.rabbit.Save(movement)
	if err != nil {
		return err
	}

	return nil
}
