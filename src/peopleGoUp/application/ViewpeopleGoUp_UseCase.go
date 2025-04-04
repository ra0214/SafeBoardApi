package application

import (
	"apiMulti/src/peopleGoUp/domain"
)

type ViewPeopleGoUp struct {
	repo domain.IPeopleGoUp
}

func NewViewPeopleGoUp(repo domain.IPeopleGoUp) *ViewPeopleGoUp {
	return &ViewPeopleGoUp{repo: repo}
}

func (vu *ViewPeopleGoUp) Execute() ([]domain.PeopleGoUp, error) {
	return vu.repo.GetAll()
}
