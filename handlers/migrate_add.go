package handlers

import (
	"log"

	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/database"
)

func HandleWithMigrateModel(cfg *config.ConfigYaml, sql []byte) string {
	db := database.ConnectionDatabase(cfg)

	err := db.MustExec(string(sql))

	if err != nil {
		log.Fatalf("%v", err)
	}

	db.Close()

	return string(sql)
}
