package infraestructure

import (
	"apiMulti/src/config"
	"apiMulti/src/movement/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IMovement = (*MySQL)(nil)

func NewMySQL() domain.IMovement {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SaveMovement(esp32_id string, aceleracion float64) error {
	query := "INSERT INTO movement (esp32_id, aceleracion) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, esp32_id, aceleracion)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Movimiento guardado correctamente: Esp32ID:%s Aceleracion:%f", esp32_id, aceleracion)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.Movement, error) {
	query := "SELECT id, esp32_id, aceleracion FROM movement"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var movementt []domain.Movement

	for rows.Next() {
		var movement domain.Movement
		if err := rows.Scan(&movement.ID, &movement.Esp32ID, &movement.Aceleracion); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}
		movementt = append(movementt, movement)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
	}
	return movementt, nil
}
