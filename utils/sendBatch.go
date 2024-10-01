package utils

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func SendBatchOverWebSocket(conn *websocket.Conn, batch Batch, mu *sync.Mutex) error {
	mu.Lock()
	defer mu.Unlock()

	err := conn.WriteJSON(batch)
	if err != nil {
		log.Printf("Error sending batch: %v", err)
		return err
	}
	return nil
}
