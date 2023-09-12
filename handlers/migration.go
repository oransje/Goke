package handlers

import (
	"fmt"
	"os"

	"github.com/vsantos1/Goke/utils"
)

func CreateMigrationsDir() (*os.File, error) {
	dir, err := os.Open("./migrations")

	if os.IsNotExist(err) {
		os.Mkdir("./migrations", os.ModePerm)
	}

	utils.HandleError(err, "[ERROR]: creating migrations directory")

	return dir, nil
}

func CreateMigrationsHistory(schema string, b []byte) {

	var FILE_PATH = fmt.Sprintf("./migrations/%s_ADD_%s.sql", utils.FormatDate(), schema)
	errs := os.WriteFile(FILE_PATH, b, os.ModePerm)

	utils.HandleError(errs, "Failed while creating, maybe forgot goke --init ?")

}
