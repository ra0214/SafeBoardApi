package domain

type IPeopleGoUpRabbitqm interface {
	Save(PeopleGoUp *PeopleGoUp) error
}