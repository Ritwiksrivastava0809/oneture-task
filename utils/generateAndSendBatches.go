package utils

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

func GenerateandSendBatches(conn *websocket.Conn, totalRecords int, recordsPerBatch int, maxRequestsPerSec int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	batch := Batch{}

	for i := 1; i <= totalRecords; i++ {
		batch.Records = append(batch.Records, Record{ID: i, Data: fmt.Sprintf("Record-%d", i)})

		if len(batch.Records) == recordsPerBatch {
			wg.Add(1)
			go func(b Batch) {
				defer wg.Done()
				if err := SendBatchOverWebSocket(conn, b, &mu); err != nil {
					log.Printf("Failed to send batch: %v", err)
				}
			}(batch)

			batch.Records = nil
			time.Sleep(time.Second / time.Duration(maxRequestsPerSec))
		}
	}

	if len(batch.Records) > 0 {
		wg.Add(1)
		go func(b Batch) {
			defer wg.Done()
			if err := SendBatchOverWebSocket(conn, b, &mu); err != nil {
				log.Printf("Failed to send batch: %v", err)
			}
		}(batch)
	}

	wg.Wait()
	log.Println("All batches sent successfully.")
}
