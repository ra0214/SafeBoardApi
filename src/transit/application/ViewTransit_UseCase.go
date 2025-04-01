package application

import (
	"apiMulti/src/transit/domain"
)

type ViewTransit struct {
	db domain.ITransit
}

func NewViewTransit(db domain.ITransit) *ViewTransit {
	return &ViewTransit{db: db}
}

func (vt *ViewTransit) Execute() ([]domain.Transit, error) {
	return vt.db.GetAll()
}