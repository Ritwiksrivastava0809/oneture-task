package utils

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

func GenerateAndSendBatchesOverWebSocket(totalRecords int, recordsPerBatch int, maxRequestsPerSec int, serverURL string) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxRequestsPerSec)

	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket server: %v", err)
	}
	defer conn.Close()

	batch := Batch{}
	for i := 1; i <= totalRecords; i++ {
		batch.Records = append(batch.Records, Record{ID: i, Data: fmt.Sprintf("Record-%d", i)})

		if len(batch.Records) == recordsPerBatch {
			sem <- struct{}{}
			wg.Add(1)

			batchToSend := batch
			go func(b Batch) {
				defer wg.Done()
				defer func() { <-sem }()

				SendBatchOverWebSocket(conn, b)
			}(batchToSend)

			batch.Records = nil
		}
	}

	if len(batch.Records) > 0 {
		sem <- struct{}{}
		wg.Add(1)

		batchToSend := batch
		go func(b Batch) {
			defer wg.Done()
			defer func() { <-sem }()

			SendBatchOverWebSocket(conn, b)
		}(batchToSend)
	}

	wg.Wait()
	log.Println("All records sent successfully, closing connection...")
}
