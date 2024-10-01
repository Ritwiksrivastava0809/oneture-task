package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-file/utils"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		var batch utils.Batch
		line := scanner.Text()

		if err := json.Unmarshal([]byte(line), &batch); err != nil {
			log.Printf("Error unmarshalling data: %v", err)
			continue
		}

		fmt.Printf("Received batch: %+v\n", batch)
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading from connection: %v", err)
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listen.Close()

	log.Println("Server is listening on port 8080...")

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
