package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/handlers"
)

var schema string
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate your SQL model from goke to database",
	Run:   migrateCmdRunner,
}

func migrateCmdRunner(cmd *cobra.Command, args []string) {

	c, err := config.ReadConfigFile("goke-config.yaml")
	if err != nil {
		panic("[Error] reading configuration file")
	}

	typ := handlers.HandleWithMigrateModel(c, handlers.ReadMigrationFileFromBytes(), schema)
	if typ.Problem != nil {
		log.Fatalf("[Error]: %v", typ.Problem)
	}

	handlers.CreateMigrationsHistory(schema, handlers.ReadMigrationFileFromBytes())

	println(typ.Message)
}
