package domain

type IMovement interface {
	SaveMovement(esp32ID string, aceleracion float64) error
	GetAll() ([]Movement, error)
}

type Movement struct {
	ID          int32   `json:"id"`
	Esp32ID     string  `json:"esp32_id"`
	Aceleracion float64 `json:"aceleracion"`
}

func NewMovement(esp32ID string, aceleracion float64) *Movement {
	return &Movement{
		Esp32ID:     esp32ID,
		Aceleracion: aceleracion,
	}
}

func (t *Movement) SetAceleracion(aceleracion float64) {
	t.Aceleracion = aceleracion
}