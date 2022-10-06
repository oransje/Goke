package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/database"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate your sql model from dummy to database",
	Run: func(cmd *cobra.Command, args []string) {
		time.Sleep(1000)

		c, err := config.ReadConfigFile("config.yaml")
		mig := database.Migrate(c, database.ReadDummyBytes())
		if err != nil {
			panic("Error while reading file config.yaml ")
		}
		println(mig)
	},
}

func init() {

	RootCmd.AddCommand(migrateCmd)
}
