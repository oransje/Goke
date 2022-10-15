package handlers

import (
	"fmt"
	"os"

	"github.com/vsantos1/Goke/utils"
)

func DropSchemaHistory(table string) {

	var DROP = fmt.Sprintf("/* DROPPED TABLE %s at %s */", table, utils.FormatDate())
	var history = []byte(DROP)
	var FILE_PATH = fmt.Sprintf("./migrations/%s_DROP_%s.sql", utils.FormatDate(),table)
	errs := os.WriteFile(FILE_PATH, history, os.ModePerm)

	if errs != nil {
		panic("Failed while creating, maybe forgot goke --init ?")
	}

}
