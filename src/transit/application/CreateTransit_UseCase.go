package application

import (
	"apiMulti/src/transit/domain"
)

type CreateTransit struct {
	rabbit domain.ITransitRabbitqm
	db domain.ITransit
}

func NewCreateTransit(r domain.ITransitRabbitqm, db domain.ITransit) *CreateTransit {
	return &CreateTransit{rabbit: r, db: db}
}

func (ct *CreateTransit) Execute(cantidad int32, tiempo string, fecha string) error {
	
	err := ct.db.SaveTransit(cantidad, tiempo, fecha)
	if err != nil {
		return err
	}

	transit := domain.NewTransit(cantidad, tiempo, fecha)

	err = ct.rabbit.Save(transit)
	if err != nil {
		return err
	}

	return nil
}