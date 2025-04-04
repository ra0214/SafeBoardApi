package application

import (
	"apiMulti/src/peopleGoDown/domain"
)

type CreatePeopleGoDown struct {
	rabbit domain.IPeopleGoDownRabbitqm
	db     domain.IPeopleGoDown
}

func NewCreatePeopleGoDown(r domain.IPeopleGoDownRabbitqm, db domain.IPeopleGoDown) *CreatePeopleGoDown {
	return &CreatePeopleGoDown{rabbit: r, db: db}
}

func (ct *CreatePeopleGoDown) Execute(esp32_id string, conteo int32) error {
	err := ct.db.SavePeopleGoDown(esp32_id, conteo)
	if err != nil {
		return err
	}

	peopleGoDown := domain.NewPeopleGoDown(esp32_id, conteo)

	err = ct.rabbit.Save(peopleGoDown)
	if err != nil {
		return err
	}

	return nil
}
