package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/database"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test your database connection from goke-config.yaml",
	Run:   testCmdRunner,
}

func testCmdRunner(cmd *cobra.Command, args []string) {
	c, err := config.ReadConfigFile("goke-config.yaml")
	if err != nil {
		panic("[Error]: reading file configuration file")
	}

	println(database.TestDatabaseconnection(c))

}
