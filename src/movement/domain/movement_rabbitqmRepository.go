package domain

type IMovementRabbitqm interface {
	Save(Movement *Movement) error
}