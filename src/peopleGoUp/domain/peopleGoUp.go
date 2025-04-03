package domain

type IPeopleGoUp interface {
	SavePeopleGoUp(esp32ID string, cantidad int32) error
	GetAll() ([]PeopleGoUp, error)
}

type PeopleGoUp struct {
	ID          int32 `json:"id"`
	Esp32ID     string `json:"esp32_id"`
	Cantidad	int32 `json:"cantidad"`
}

func NewPeopleGoUp(esp32ID string ,cantidad int32) *PeopleGoUp {
	return &PeopleGoUp{
		Esp32ID: esp32ID, 
		Cantidad: cantidad}
}

func (t *PeopleGoUp) SetCantidad(cantidad int32) {
	t.Cantidad = cantidad
}