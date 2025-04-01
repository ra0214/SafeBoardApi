package domain

type ITransitRabbitqm interface {
	Save(Transit *Transit) error
}