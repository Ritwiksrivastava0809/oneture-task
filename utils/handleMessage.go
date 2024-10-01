package utils

import (
	"log"

	"github.com/gorilla/websocket"
)

func HandleMessage(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Printf("Received batch: %s", message)

		if err := conn.WriteMessage(websocket.TextMessage, []byte("Batch received")); err != nil {
			log.Printf("Error sending acknowledgment: %v", err)
			break
		}
	}
}
