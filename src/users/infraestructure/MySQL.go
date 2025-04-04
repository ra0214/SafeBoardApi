package infraestructure

import (
	"apiMulti/src/config"
	"apiMulti/src/users/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IUser = (*MySQL)(nil)

func NewMySQL() domain.IUser {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SaveUser(userName string, email string, password string, esp32ID string) error {
	query := "INSERT INTO users (userName, email, password, esp32_id) VALUES (?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, userName, email, password, esp32ID)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Usuario creado correctamente: Username:%s Email:%s ESP32ID:%s", userName, email, esp32ID)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.User, error) {
	query := "SELECT id, userName, email, password, esp32_id FROM users"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.ESP32ID); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
	}
	return users, nil
}

func (mysql *MySQL) UpdateUser(id int32, userName string, email string, password string, esp32ID string) error {
	query := "UPDATE users SET userName = ?, email = ?, password = ?, esp32_id = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, userName, email, password, esp32ID, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Usuario actualizado correctamente: ID: %d Username:%s Email: %s ESP32ID: %s", id, userName, email, esp32ID)
	} else {
		log.Println("[MySQL] - No se actualizó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) DeleteUser(id int32) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Usuario eliminado correctamente: ID: %d", id)
	} else {
		log.Println("[MySQL] - No se eliminó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetUserByCredentials(userName string) (*domain.User, error) {
	query := "SELECT id, userName, email, password, esp32_id FROM users WHERE userName = ?"
	row, err := mysql.conn.FetchRow(query, userName)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	var user domain.User
	err = row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.ESP32ID)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado")
	}

	return &user, nil
}

func (mysql *MySQL) GetUserByESP32ID(esp32ID string) (*domain.User, error) {
	query := "SELECT id, userName, email, password, esp32_id FROM users WHERE esp32_id = ?"
	row, err := mysql.conn.FetchRow(query, esp32ID)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	var user domain.User
	err = row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.ESP32ID)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado para ESP32ID: %s", esp32ID)
	}

	return &user, nil
}
