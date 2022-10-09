package database

import (
	"bufio"
	"log"
	"os"
)

// FIX
func readDummyByte() []byte {
	file, err := os.OpenFile("dummy.sql", os.O_RDONLY, os.ModePerm)
	var byt []byte
	if err != nil {
		log.Fatalf("failed while trying to %v", err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		byt = sc.Bytes()
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	return byt

}
