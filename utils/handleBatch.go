package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleBatch(c *gin.Context) {
	var batch Batch

	if err := c.ShouldBindJSON(&batch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	for _, record := range batch.Records {
		log.Printf("Processing record ID: %d, Data: %s", record.ID, record.Data)
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Batch processed",
		"record_count": len(batch.Records),
	})
}
