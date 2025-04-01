package infraestructure

import (
	"apiMulti/src/config"
	"apiMulti/src/transit/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.ITransit = (*MySQL)(nil)

func NewMySQL() domain.ITransit {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SaveTransit(cantidad int32, tiempo string, fecha string) error {
	query := "INSERT INTO transit (cantidad, tiempo, fecha) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, cantidad, tiempo, fecha)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Transito de personas guardado correctamente: Cantidad:%s Tiempo:%s Fecha:%s", cantidad, tiempo, fecha)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.Transit, error) {
    query := "SELECT id, cantidad, tiempo, fecha FROM transit"
    rows, err := mysql.conn.FetchRows(query)
    if err != nil {
        return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
    }
    defer rows.Close()

    var transitt []domain.Transit

    for rows.Next() {
        var transit domain.Transit
        if err := rows.Scan(&transit.ID, &transit.Cantidad, &transit.Tiempo, &transit.Fecha); err != nil {
            return nil, fmt.Errorf("Error al escanear la fila: %v", err)
        }
        transitt = append(transitt, transit)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
    }
    return transitt, nil
}