package infraestructure

import (
	"apiMulti/src/config"
	"apiMulti/src/peopleGoUp/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IPeopleGoUp = (*MySQL)(nil)

func NewMySQL() domain.IPeopleGoUp {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SavePeopleGoUp(esp32_id string, cantidad int32) error {
	query := "INSERT INTO goup (esp32_id, conteo) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, esp32_id, cantidad)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Conteo de personas guardado correctamente: Esp32ID:%s Cantidad:%s ", esp32_id, cantidad)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.PeopleGoUp, error) {
	query := "SELECT id, esp32_id, conteo FROM goup ORDER BY id DESC"
	log.Printf("[MySQL] Ejecutando query: %s", query)

	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		log.Printf("[MySQL] Error en FetchRows: %v", err)
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var peopleGoUpp []domain.PeopleGoUp

	for rows.Next() {
		var peopleGoUp domain.PeopleGoUp
		var esp32ID string
		var conteo int32

		if err := rows.Scan(&peopleGoUp.ID, &esp32ID, &conteo); err != nil {
			log.Printf("[MySQL] Error al escanear fila: %v", err)
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}

		peopleGoUp.Esp32ID = esp32ID
		peopleGoUp.Conteo = conteo

		log.Printf("[MySQL] Fila escaneada: ID=%d, Esp32ID='%s', Conteo=%d",
			peopleGoUp.ID, peopleGoUp.Esp32ID, peopleGoUp.Conteo)

		peopleGoUpp = append(peopleGoUpp, peopleGoUp)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[MySQL] Error al iterar filas: %v", err)
		return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
	}

	log.Printf("[MySQL] Total de registros encontrados: %d", len(peopleGoUpp))
	return peopleGoUpp, nil
}
