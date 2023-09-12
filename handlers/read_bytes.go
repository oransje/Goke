package handlers

import (
	"os"

	"github.com/vsantos1/Goke/utils"
)

// TODO: WHY NOT CREATE A MIGRATION SIMPLE LEXER AND PARSER?
func ReadMigrationFileFromBytes() []byte {
	file, err := os.ReadFile("goke.sql")
	utils.HandleError(err, "[ERROR] while reading goke.sql file")
	return file
}
