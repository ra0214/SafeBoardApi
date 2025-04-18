package infraestructure

import (
	"log"
	"apiMulti/src/config"
)

func Init() {
	mysqlRepo := NewMySQL()

	rabbitMQRepo, err := config.GetChannel()
	if err != nil {
		log.Fatalf("Error al inicializar RabbitMQ: %v", err)
	}
	defer rabbitMQRepo.Close()

	router := SetupRouter(mysqlRepo)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
