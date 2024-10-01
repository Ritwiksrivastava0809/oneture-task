package main

import (
	"go-file/utils"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	serverURL := "ws://localhost:8080/ws"
	totalRecords := 100000
	recordsPerBatch := 100
	maxRequestsPerSec := 1000

	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()

	startTime := time.Now()

	utils.GenerateandSendBatches(conn, totalRecords, recordsPerBatch, maxRequestsPerSec)

	endTime := time.Now()

	duration := endTime.Sub(startTime)

	log.Printf("All batches sent successfully in %s.", duration)
	log.Printf("Total records sent: %d", totalRecords)
	log.Printf("Records per batch: %d", recordsPerBatch)
	log.Printf("Max requests per second: %d", maxRequestsPerSec)
}
