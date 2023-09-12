package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/database"
	"github.com/vsantos1/Goke/utils"
)

type ReturnType struct {
	Message string
	Problem error
}

func DropTableQuery(c *config.ConfigYaml, table string) ReturnType {
	db := database.ConnectionDatabase(c)
	if table == "" {
		log.Fatal("Cannot delete empty table name")

	}

	defer db.Close()

	var QUERY_DROP = fmt.Sprintf("DROP TABLE %s", table)

	_, err := db.Exec(QUERY_DROP)

	utils.HandleError(err, "Error while dropping table")

	now := time.Now()

	var msg = fmt.Sprintf("Successfully dropped Table %s at %s", table, now.Format(time.RFC822))
	var returnType = ReturnType{
		Message: msg,
		Problem: nil,
	}

	return returnType
}
