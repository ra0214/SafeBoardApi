package domain

type ITransit interface {
	SaveTransit(cantidad int32, tiempo string, fecha string) error
	GetAll() ([]Transit, error)
}

type Transit struct {
	ID          int32 `json:"id"`
	Cantidad	int32 `json:"cantidad"`
	Tiempo 		string `json:"tiempo"`
	Fecha		string `json:"fecha"`
}

func NewTransit(cantidad int32, tiempo, fecha string) *Transit {
	return &Transit{ Cantidad: cantidad, Tiempo: tiempo, Fecha: fecha}
}

func (t *Transit) SetCantidad(cantidad int32) {
	t.Cantidad = cantidad
}