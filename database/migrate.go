package database

import (
	"bufio"
	"log"
	"os"

	"github.com/vsantos1/Goke/config"
)

// FIX
func readDummyByte() []byte {
	file, err := os.OpenFile("dummy.sql", os.O_RDONLY, os.ModePerm)
	var t []byte
	if err != nil {
		log.Fatalf("failed while trying to %v", err)
	}

	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		t = sc.Bytes()
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
	}

	return t

}

func Migrate(cfg *config.ConfigYaml, sql []byte) string {
	db := ConnectionDatabase(cfg)

	err := db.MustExec(string(sql))

	if err != nil {
		log.Fatalf("%v", err)
	}
	db.Close()
	return string(sql)
}
