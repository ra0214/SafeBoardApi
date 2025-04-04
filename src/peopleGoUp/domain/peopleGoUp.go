package domain

type IPeopleGoUp interface {
	SavePeopleGoUp(esp32ID string, cantidad int32) error
	GetAll() ([]PeopleGoUp, error)
}

type PeopleGoUp struct {
	ID      int32  `json:"id"`
	Esp32ID string `json:"esp32_id"`
	Conteo  int32  `json:"conteo"`
}

func NewPeopleGoUp(esp32ID string, conteo int32) *PeopleGoUp {
	return &PeopleGoUp{
		Esp32ID: esp32ID,
		Conteo:  conteo,
	}
}

func (t *PeopleGoUp) SetCantidad(conteo int32) {
	t.Conteo = conteo
}
