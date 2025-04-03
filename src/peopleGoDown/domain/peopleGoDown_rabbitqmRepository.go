package domain

type IPeopleGoDownRabbitqm interface {
	Save(PeopleGoDown *PeopleGoDown) error
}