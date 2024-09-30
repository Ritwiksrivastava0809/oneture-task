package main

import (
	"github.com/gin-gonic/gin"
	"go-file/utils"
	"log"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()

	router.POST("/process-batch", utils.HandleBatch)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Gin server is running on port 8080")
	log.Fatal(s.ListenAndServe())
}
