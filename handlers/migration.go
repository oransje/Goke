package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/vsantos1/Goke/utils"
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

func CreateHistoryDb(schema string, b []byte) {

	var FILE_PATH = fmt.Sprintf("./migrations/%s_ADD_%s.sql", schema, utils.FormatDate())
	errs := os.WriteFile(FILE_PATH, b, os.ModePerm)

	if errs != nil {
		panic("Failed while creating, maybe forgot goke --init ?")
	}

}
