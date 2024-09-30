package main

import (
	"go-file/utils" // Adjust the import path as necessary
	"log"
	"time"
)

const (
	ServerURL         = "http://localhost:8080/process-batch"
	RecordsPerBatch   = 10
	TotalRecords      = 100000 // Simulated total records
	MaxRequestsPerSec = 1000   // Max client requests per second
)

func main() {
	log.Println("Client starting to generate and send records...")

	start := time.Now()

	// Generate and send batches of records
	utils.GenerateAndSendBatches(TotalRecords, RecordsPerBatch, MaxRequestsPerSec, ServerURL)

	log.Printf("All records sent in %v", time.Since(start))
}
