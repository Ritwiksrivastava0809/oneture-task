package main

import (
	"fmt"
	"go-file/utils"
	"log"
	"net"
	"time"
)

func main() {
	serverAddr := "localhost:8080"
	totalRecords := 100000
	recordsPerBatch := 100
	maxRequestsPerSec := 1000

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	startTime := time.Now()
	batch := utils.Batch{}

	for i := 1; i <= totalRecords; i++ {
		batch.Records = append(batch.Records, utils.Record{ID: i, Data: fmt.Sprintf("Record-%d", i)})

		if len(batch.Records) == recordsPerBatch {
			if err := utils.SendBatch(conn, batch); err != nil {
				log.Printf("Failed to send batch: %v", err)
			}
			batch.Records = nil
			time.Sleep(time.Second / time.Duration(maxRequestsPerSec))
		}
	}

	if len(batch.Records) > 0 {
		if err := utils.SendBatch(conn, batch); err != nil {
			log.Printf("Failed to send final batch: %v", err)
		}
	}

	elapsedTime := time.Since(startTime)
	log.Printf("All batches sent successfully in %s.", elapsedTime)
	log.Printf("Total records sent: %d", totalRecords)
	log.Printf("Records per batch: %d", recordsPerBatch)
	log.Printf("Max requests per second: %d", maxRequestsPerSec)
}
