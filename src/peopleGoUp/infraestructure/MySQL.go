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

// Constructor de la conexión a MySQL
func NewMySQL() domain.IPeopleGoUp {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Guardar el conteo de personas en la base de datos
func (mysql *MySQL) SavePeopleGoUp(esp32ID string, cantidad int32) error {
	query := "INSERT INTO goup (esp32_id, conteo) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, esp32ID, cantidad)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener las filas afectadas: %v", err)
	}

	if rowsAffected == 1 {
		log.Printf("[MySQL] - Conteo de personas guardado correctamente: Esp32ID:%s Cantidad:%d", esp32ID, cantidad)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

// Obtener todos los registros de conteo de personas
func (mysql *MySQL) GetAll() ([]domain.PeopleGoUp, error) {
	query := "SELECT id, esp32_id, conteo FROM goup"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var peopleGoUpList []domain.PeopleGoUp

	for rows.Next() {
		var peopleGoUp domain.PeopleGoUp
		if err := rows.Scan(&peopleGoUp.ID, &peopleGoUp.Esp32ID, &peopleGoUp.Conteo); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		peopleGoUpList = append(peopleGoUpList, peopleGoUp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}
	return peopleGoUpList, nil
}
