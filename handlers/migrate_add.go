package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/database"
)

func HandleWithMigrateModel(cfg *config.ConfigYaml, sql []byte, file string) ReturnType {
	db := database.ConnectionDatabase(cfg)

	_,err := db.Exec(string(sql))
	if err != nil {
		log.Fatalf("Error: %v",err)
		return ReturnType{
			Message: fmt.Sprintf("Error: %v",err),
			Problem: err,
		}
	}
	
	defer db.Close()
	now := time.Now()
	var msg = fmt.Sprintf("successfully created table file name start's with %s at %s", file, now.Format(time.RFC822))
	var returnType = ReturnType{
		Message: msg,
		Problem: nil,
	}
	return returnType
}
