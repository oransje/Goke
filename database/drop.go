package database

import (
	"fmt"
	"log"

	"github.com/vsantos1/Goke/config"
)

func DropDatabaseSchema(c *config.ConfigYaml, table string) (string, error) {
	db := ConnectionDatabase(c)
	var QUERY_DROP = fmt.Sprintf("DROP TABLE IF EXISTS %s", table)

	err := db.MustExec(QUERY_DROP)

	if err != nil {
		log.Fatalf("%v", err)
	}

	db.Close()
	return "Deleted", nil
}
