package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/database"
)

var table string

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop database model from database",
	Run: func(cmd *cobra.Command, args []string) {
		time.Sleep(1000)
		c, err := config.ReadConfigFile("config.yaml")
		if err != nil {
			panic("Error while reading file config.yaml ")
		}
		mig, err := database.DropDatabaseSchema(c, table)
		if err != nil {
			panic("Failed to drop table ")

		}

		println(mig)
	},
}

func init() {
	dropCmd.Flags().StringVarP(&table, "table", "t", "", "Drop table from database if exists")
	RootCmd.AddCommand(dropCmd)
}
