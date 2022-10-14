package database

import (
	"fmt"
	"log"

	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/utils"
)

func DropDatabaseSchema(c *config.ConfigYaml, table string) (string, error) {
	db := ConnectionDatabase(c)
	if table == "" {
		log.Fatal("cannot delete empty table name")

	}

	var QUERY_DROP = fmt.Sprintf("DROP TABLE %s", table)

	db.MustExec(QUERY_DROP)
	var message = fmt.Sprintf("successfully dropped database %s at %s", table, utils.FormatDate())

	db.Close()
	return message, nil
}
