package infraestructure

import (
	"apiMulti/src/config"
	"log"
)

func InitTransit() error {
	log.Println("Inicializando datos...")

	db, err := config.GetDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("Conexión a la base de datos para citas establecida correctamente")
	return nil
}