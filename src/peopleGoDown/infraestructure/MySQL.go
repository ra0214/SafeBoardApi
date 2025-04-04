package infraestructure

import (
	"apiMulti/src/config"
	"apiMulti/src/peopleGoDown/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IPeopleGoDown = (*MySQL)(nil)

func NewMySQL() domain.IPeopleGoDown {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SavePeopleGoDown(esp32_id string, cantidad int32) error {
	query := "INSERT INTO godown (esp32_id, conteo) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, esp32_id, cantidad)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Conteo de personas guardado correctamente: Cantidad:%s ", cantidad)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.PeopleGoDown, error) {
	query := "SELECT id, esp32_id, conteo FROM godown ORDER BY id DESC"
	log.Printf("[MySQL] Ejecutando query: %s", query)

	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		log.Printf("[MySQL] Error en FetchRows: %v", err)
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var peopleGoDowns []domain.PeopleGoDown

	for rows.Next() {
		var p domain.PeopleGoDown
		if err := rows.Scan(&p.ID, &p.ESP32ID, &p.Conteo); err != nil {
			log.Printf("[MySQL] Error al escanear fila: %v", err)
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}
		peopleGoDowns = append(peopleGoDowns, p)
	}

	log.Printf("[MySQL] Registros encontrados: %d", len(peopleGoDowns))
	return peopleGoDowns, nil
}
