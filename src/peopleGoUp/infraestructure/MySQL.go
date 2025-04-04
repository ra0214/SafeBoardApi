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
	// Simplificamos la consulta para depuración
	query := "SELECT id, esp32_id, conteo FROM goup"
	log.Printf("[MySQL] Query: %s", query)

	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		log.Printf("[MySQL] Error FetchRows: %v", err)
		return nil, err
	}
	defer rows.Close()

	var peopleGoUpp []domain.PeopleGoUp

	// Debug: Imprimir columnas
	cols, _ := rows.Columns()
	log.Printf("[MySQL] Columnas: %v", cols)

	for rows.Next() {
		var p domain.PeopleGoUp

		// Crear slice de interfaces para el escaneo
		values := make([]interface{}, 3)
		values[0] = &p.ID
		values[1] = &p.Esp32ID
		values[2] = &p.Conteo

		if err := rows.Scan(values...); err != nil {
			log.Printf("[MySQL] Error Scan: %v", err)
			return nil, err
		}

		log.Printf("[MySQL] Datos escaneados: %+v", p)
		peopleGoUpp = append(peopleGoUpp, p)
	}

	return peopleGoUpp, nil
}
