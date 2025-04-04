package application

import (
	"apiMulti/src/peopleGoDown/domain"
	"log"
)

type CreatePeopleGoDown struct {
	rabbitRepo domain.IPeopleGoDownRabbitqm
	mysqlRepo  domain.IPeopleGoDown
}

func NewCreatePeopleGoDown(rabbitRepo domain.IPeopleGoDownRabbitqm, mysqlRepo domain.IPeopleGoDown) *CreatePeopleGoDown {
	return &CreatePeopleGoDown{
		rabbitRepo: rabbitRepo,
		mysqlRepo:  mysqlRepo,
	}
}

func (c *CreatePeopleGoDown) Execute(esp32_id string, cantidad int32) error {
	log.Printf("[UseCase] Iniciando creación de PeopleGoDown: ESP32_ID=%s, Cantidad=%d", esp32_id, cantidad)

	// Guardar en MySQL
	err := c.mysqlRepo.SavePeopleGoDown(esp32_id, cantidad)
	if err != nil {
		log.Printf("[UseCase] Error al guardar en MySQL: %v", err)
		return err
	}

	// Publicar en RabbitMQ


	log.Printf("[UseCase] PeopleGoDown creado exitosamente")
	return nil
}
