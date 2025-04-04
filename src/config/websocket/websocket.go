package websocket

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Permite todas las conexiones en desarrollo
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func HandleWebSocket(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error al establecer conexión WebSocket: %v", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			delete(clients, ws)
			break
		}
	}
}

func BroadcastMessage(messageType string, data interface{}) {
	message := Message{
		Type: messageType,
		Data: data,
	}
	broadcast <- message
}

func init() {
	go handleMessages()
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error al enviar mensaje: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
