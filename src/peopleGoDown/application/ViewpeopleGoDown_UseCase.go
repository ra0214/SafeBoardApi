package application

import (
	"apiMulti/src/peopleGoDown/domain"
)

type ViewPeopleGoDown struct {
	repo domain.IPeopleGoDown
}

func NewViewPeopleGoDown(repo domain.IPeopleGoDown) *ViewPeopleGoDown {
	return &ViewPeopleGoDown{repo: repo}
}

func (vu *ViewPeopleGoDown) Execute() ([]domain.PeopleGoDown, error) {
	return vu.repo.GetAll()
}
