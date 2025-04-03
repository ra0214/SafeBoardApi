package domain

type IPeopleGoUp interface {
	SavePeopleGoUp(cantidad int32) error
	GetAll() ([]PeopleGoUp, error)
}

type PeopleGoUp struct {
	ID          int32 `json:"id"`
	Cantidad	int32 `json:"cantidad"`
}

func NewPeopleGoUp(cantidad int32) *PeopleGoUp {
	return &PeopleGoUp{ Cantidad: cantidad}
}

func (t *PeopleGoUp) SetCantidad(cantidad int32) {
	t.Cantidad = cantidad
}