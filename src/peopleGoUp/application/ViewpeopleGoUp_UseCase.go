package application

import (
	"apiMulti/src/peopleGoUp/domain"
	"log"
)

type ViewPeopleGoUp struct {
	repo domain.IPeopleGoUp
}

func NewViewPeopleGoUp(repo domain.IPeopleGoUp) *ViewPeopleGoUp {
	return &ViewPeopleGoUp{repo: repo}
}

func (vu *ViewPeopleGoUp) Execute() ([]domain.PeopleGoUp, error) {
	log.Println("[UseCase] Iniciando GetAll")
	data, err := vu.repo.GetAll()
	if err != nil {
		log.Printf("[UseCase] Error en GetAll: %v", err)
		return nil, err
	}
	log.Printf("[UseCase] Datos obtenidos exitosamente: %d registros", len(data))
	return data, nil
}
