package domain

type IPeopleGoDown interface {
	SavePeopleGoDown(esp32_id string, cantidad int32) error
	GetAll() ([]PeopleGoDown, error)
}

type PeopleGoDown struct {
	ID      int32  `json:"id"`
	ESP32ID string `json:"esp32_id"`
	Conteo  int32  `json:"conteo"`
}

func NewPeopleGoDown(esp32ID string, conteo int32) *PeopleGoDown {
	return &PeopleGoDown{
		ESP32ID: esp32ID,
		Conteo:  conteo,
	}
}

func (t *PeopleGoDown) SetCantidad(conteo int32) {
	t.Conteo = conteo
}
