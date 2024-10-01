package utils

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

var writeMutex sync.Mutex

func SendBatchOverWebSocket(conn *websocket.Conn, batch Batch) {
	writeMutex.Lock()
	defer writeMutex.Unlock()

	data, err := json.Marshal(batch)
	if err != nil {
		log.Printf("Failed to serialize batch: %v", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Printf("Failed to send batch: %v", err)
		return
	}

	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Printf("Failed to read acknowledgment: %v", err)
		return
	}

	log.Printf("Received acknowledgment from server: %s", message)
}
