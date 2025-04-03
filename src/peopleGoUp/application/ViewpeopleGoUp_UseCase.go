package application

import (
	"apiMulti/src/peopleGoUp/domain"
)

type ViewPeopleGoUp struct {
	db domain.IPeopleGoUp
}

func NewViewPeopleGoUp(db domain.IPeopleGoUp) *ViewPeopleGoUp {
	return &ViewPeopleGoUp{db: db}
}

func (vt *ViewPeopleGoUp) Execute() ([]domain.PeopleGoUp, error) {
	return vt.db.GetAll()
}