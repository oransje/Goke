package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/database"
	"github.com/vsantos1/Goke/utils"
)
type JSONData struct {
	Values []float64 `json:"Values"`
	Dates []string `json:"Dates"`
  }

func DumpSchema(table string,cfg *config.ConfigYaml) ReturnType {
	db := database.ConnectionDatabase(cfg)
	defer db.Close()
	now := time.Now()
	var SELECT_QUERY = fmt.Sprintf("SELECT* FROM %s", table)

	res,_ := db.Query(SELECT_QUERY)
	
	a := utils.Jsonify(res)
	out,err := json.Marshal(a)

	if err != nil {
		panic(err)
	}
	var JSON_FILE = fmt.Sprintf("./migrations/%s_DUMP_%s.json", table, utils.FormatDate())
	errs := os.WriteFile(JSON_FILE, out, os.ModePerm)
	if errs != nil {
		panic("Failed while creating, maybe forgot goke --init ?")
	}


	var msg = fmt.Sprintf("successfully dumpped table %s at %s", table, now.Format(time.RFC822))
	
	var returnType = ReturnType{
		Message: msg,
		Problem: nil,
	}

	return returnType
}