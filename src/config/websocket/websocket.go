package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	Type       string      `json:"type"`
	Data       interface{} `json:"data"`
	TargetUser string      `json:"targetUser"` // Este campo determina a qué ESP32 va dirigido
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]string) // Mapa de conexiones a ESP32 IDs
var broadcast = make(chan Message)

func HandleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error al establecer conexión WebSocket: %v", err)
		return
	}
	defer ws.Close()

	// Esperar el mensaje de registro con el ESP32 ID
	_, msgBytes, err := ws.ReadMessage()
	if err != nil {
		log.Printf("Error al leer mensaje inicial: %v", err)
		return
	}

	var initMsg struct {
		Type     string `json:"type"`
		ID_ESP32 string `json:"id_esp32"`
	}

	if err := json.Unmarshal(msgBytes, &initMsg); err != nil {
		log.Printf("Error parseando mensaje inicial: %v", err)
		return
	}

	if initMsg.Type != "register" || initMsg.ID_ESP32 == "" {
		log.Println("Registro inválido: se esperaba mensaje de tipo 'register' con ID_ESP32")
		return
	}

	// Registrar el cliente con su ESP32 ID
	clients[ws] = initMsg.ID_ESP32
	log.Printf("Cliente registrado con ID_ESP32: %s", initMsg.ID_ESP32)

	// Mantener la conexión abierta
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Cliente desconectado: %v", initMsg.ID_ESP32)
			delete(clients, ws)
			break
		}
	}
}

func BroadcastMessage(messageType string, data interface{}, targetUser string) {
	message := Message{
		Type:       messageType,
		Data:       data,
		TargetUser: targetUser,
	}
	broadcast <- message
}

func init() {
	go handleMessages()
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client, clientID := range clients {
			// Si el mensaje tiene un TargetUser, solo enviarlo a ese cliente
			if msg.TargetUser != "" && msg.TargetUser != clientID {
				continue
			}
			
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error al enviar mensaje: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}