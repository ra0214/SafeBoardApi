package application

import (
	"apiMulti/src/peopleGoUp/domain"
)

type CreatepeopleGoUp struct {
	rabbit domain.IPeopleGoUpRabbitqm
	db domain.IPeopleGoUp
}

func NewCreatepeopleGoUp(r domain.IPeopleGoUpRabbitqm, db domain.IPeopleGoUp) *CreatepeopleGoUp {
	return &CreatepeopleGoUp{rabbit: r, db: db}
}

func (ct *CreatepeopleGoUp) Execute(esp32_id string ,conteo int32) error {
	
	err := ct.db.SavePeopleGoUp(esp32_id, conteo)
	if err != nil {
		return err
	}

	peopleGoUp := domain.NewPeopleGoUp(esp32_id, conteo)

	err = ct.rabbit.Save(peopleGoUp)
	if err != nil {
		return err
	}

	return nil
}