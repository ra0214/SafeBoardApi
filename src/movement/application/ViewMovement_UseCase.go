package application

import (
	"apiMulti/src/movement/domain"
)

type ViewMovement struct {
	db domain.IMovement
}

func NewViewMovement(db domain.IMovement) *ViewMovement {
	return &ViewMovement{db: db}
}

func (vt ViewMovement) Execute() ([]domain.Movement, error) {
	return vt.db.GetAll()
}