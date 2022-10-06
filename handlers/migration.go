package handlers

import (
	"log"
	"os"
)

func CreateMigrationsDir() (*os.File, error) {
	dir, err := os.Open("./migrations")

	if os.IsNotExist(err) {
		os.Mkdir("./migrations", os.ModePerm)
	}

	if dir != nil {
		log.Fatalf("Already initialized, check %v folder", dir.Name())
	}

	return dir, nil
}
