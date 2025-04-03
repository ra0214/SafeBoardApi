package application

import (
	"apiMulti/src/peopleGoDown/domain"
)

type ViewPeopleGoDown struct {
	db domain.IPeopleGoDown
}

func NewViewPeopleGoDown(db domain.IPeopleGoDown) *ViewPeopleGoDown {
	return &ViewPeopleGoDown{db: db}
}

func (vt *ViewPeopleGoDown) Execute() ([]domain.PeopleGoDown, error) {
	return vt.db.GetAll()
}