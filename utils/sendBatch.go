package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// SendBatch sends a batch of records to the specified server URL
func SendBatch(batch Batch, serverURL string) {
	jsonData, err := json.Marshal(batch)
	if err != nil {
		log.Printf("Failed to marshal batch: %v", err)
		return
	}

	resp, err := http.Post(serverURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to send batch: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read response body to get more details about the error
		var responseBody bytes.Buffer
		if _, err := responseBody.ReadFrom(resp.Body); err == nil {
			log.Printf("Server responded with status: %d, body: %s", resp.StatusCode, responseBody.String())
		} else {
			log.Printf("Server responded with status: %d, but failed to read body: %v", resp.StatusCode, err)
		}
	} else {
		log.Printf("Batch of %d records sent successfully", len(batch.Records))
	}
}
