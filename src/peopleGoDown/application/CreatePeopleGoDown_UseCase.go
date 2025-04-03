package application

import (
	"apiMulti/src/peopleGoDown/domain"
)

type CreatepeopleGoDown struct {
	rabbit domain.IPeopleGoDownRabbitqm
	db domain.IPeopleGoDown
}

func NewCreatepeopleGoDown(r domain.IPeopleGoDownRabbitqm, db domain.IPeopleGoDown) *CreatepeopleGoDown {
	return &CreatepeopleGoDown{rabbit: r, db: db}
}

func (ct *CreatepeopleGoDown) Execute(esp32_id string ,conteo int32) error {
	
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