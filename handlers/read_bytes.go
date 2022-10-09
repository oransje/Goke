package handlers

import (
	"log"
	"os"
)

func ReadDummyBytes() []byte {
	file, err := os.ReadFile("dummy.sql")
	if err != nil {
		log.Fatalf("failed while trying to %v", err)
	}

	return file
}
