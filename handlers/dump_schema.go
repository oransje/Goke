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

func DumpSchema(table string, cfg *config.ConfigYaml) ReturnType {
	db := database.ConnectionDatabase(cfg)
	defer db.Close()
	now := time.Now()
	var SELECT_QUERY = fmt.Sprintf("SELECT * FROM %s", table)

	res, err := db.Query(SELECT_QUERY)
	utils.HandleError(err, fmt.Sprintf("[ERROR]: dumping table %s", table))
	jsonified := utils.Jsonify(res)

	out, err := json.MarshalIndent(jsonified, "", "  ")

	utils.HandleError(err, "Error Marshaling to json")

	var JSON_FILE = fmt.Sprintf("./migrations/%s_DUMP_%s.json", table, utils.FormatDate())
	errs := os.WriteFile(JSON_FILE, out, os.ModePerm)

	utils.HandleError(errs, "Error writing to the file")

	var msg = fmt.Sprintf("Successfully dumpped table %s at %s", table, now.Format(time.RFC822))

	var returnType = ReturnType{
		Message: msg,
		Problem: nil,
	}

	return returnType
}
