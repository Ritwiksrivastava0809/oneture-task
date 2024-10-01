package main

import (
	"fmt"
	"go-file/utils"
	"time" // Import the time package
)

func main() {
	serverURL := "ws://localhost:8080/ws"

	totalRecords := 100000
	recordsPerBatch := 10
	maxRequestsPerSec := 1000

	fmt.Println("Starting to send records via WebSocket...")

	// Record start time
	startTime := time.Now()

	// Generate and send batches over WebSocket
	utils.GenerateAndSendBatchesOverWebSocket(totalRecords, recordsPerBatch, maxRequestsPerSec, serverURL)

	// Record end time
	endTime := time.Now()

	// Calculate and print the duration
	duration := endTime.Sub(startTime)
	fmt.Printf("All records sent successfully in %s\n", duration)
}
