package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/database"
)

type ReturnType struct {
	Message string
	Problem error
}

func DropDatabaseSchema(c *config.ConfigYaml, table string) ReturnType {
	db := database.ConnectionDatabase(c)
	if table == "" {
		log.Fatal("cannot delete empty table name")

	}

	var QUERY_DROP = fmt.Sprintf("DROP TABLE %s", table)

	db.MustExec(QUERY_DROP)
	now := time.Now()

	var msg = fmt.Sprintf("successfully dropped database %s at %s", table, now.Format(time.RFC822))
	var returnType = ReturnType{
		Message: msg,
		Problem: nil,
	}

	db.Close()
	return returnType
}
