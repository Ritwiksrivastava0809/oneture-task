package utils

import (
	"encoding/json"
	"net"
)

func SendBatch(conn net.Conn, batch Batch) error {
	data, err := json.Marshal(batch)
	if err != nil {
		return err
	}

	_, err = conn.Write(data)
	if err != nil {
		return err
	}

	_, err = conn.Write([]byte("\n"))
	return err
}
