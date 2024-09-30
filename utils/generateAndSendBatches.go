package utils

import (
	"fmt"
	"sync"
)

func GenerateAndSendBatches(totalRecords int, recordsPerBatch int, maxRequestsPerSec int, serverURL string) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxRequestsPerSec)
	batch := Batch{}
	for i := 1; i <= totalRecords; i++ {
		batch.Records = append(batch.Records, Record{ID: i, Data: fmt.Sprintf("Record-%d", i)})

		if len(batch.Records) == recordsPerBatch {
			sem <- struct{}{}
			wg.Add(1)

			go func(b Batch) {
				defer wg.Done()
				defer func() { <-sem }()
				SendBatch(b, serverURL)
			}(batch)

			batch.Records = nil
		}
	}

	if len(batch.Records) > 0 {
		sem <- struct{}{}
		wg.Add(1)

		go func(b Batch) {
			defer wg.Done()
			defer func() { <-sem }()
			SendBatch(b, serverURL)
		}(batch)
	}

	wg.Wait()
}
