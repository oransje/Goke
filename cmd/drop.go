package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/handlers"
)

var table string

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop table from database",
	RunE: func(cmd *cobra.Command, args []string) error {
		time.Sleep(1000)
		c, err := config.ReadConfigFile("goke-config.yaml")
		if err != nil {
			panic("Error while reading file goke-config.yaml ")
		}
		t := handlers.DropDatabaseSchema(c, table)
		handlers.DropSchemaHistory(table)
		println(t.Message)
		if t.Problem != nil {
			panic("Failed to drop table ")

		}
		return nil
	},
}

func init() {
	dropCmd.Flags().StringVarP(&table, "table", "t", "", "Drop table from database if exists")
	RootCmd.AddCommand(dropCmd)
}
