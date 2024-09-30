package main

import (
	"go-file/utils"
	"log"
	"time"
)

const (
	ServerURL         = "http://localhost:8080/process-batch"
	RecordsPerBatch   = 10
	TotalRecords      = 100000
	MaxRequestsPerSec = 1000
)

func main() {
	log.Println("Client starting to generate and send records...")

	start := time.Now()

	utils.GenerateAndSendBatches(TotalRecords, RecordsPerBatch, MaxRequestsPerSec, ServerURL)

	log.Printf("All records sent in %v", time.Since(start))
}
