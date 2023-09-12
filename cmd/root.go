package cmd

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/handlers"
	"github.com/vsantos1/Goke/utils"
)

var version bool
var initGoke bool
var RootCmd = &cobra.Command{
	SilenceUsage:     true,
	Short:            "Goke is a manager for database/sql",
	PersistentPreRun: rootCmdRunner,
	Run:              entryPointCmdRunner,
}

func rootCmdRunner(cmd *cobra.Command, args []string) {
	// Create base configuration to connect on database if not exists
	file, err := os.OpenFile("goke-config.yaml", 0, 0644)
	if os.IsNotExist(err) {
		file, err := os.Create("goke-config.yaml")
		dummy, errs := os.Create("goke.sql")
		if errs != nil {
			panic("Failed creating goke.sql")
		}
		dummy.WriteString(TODO_DUMMY)
		file.WriteString("username: root\n")
		file.WriteString("password: root\n")
		file.WriteString("dialect: sqlite # using SQLITE by default\n")
		file.WriteString("sslmode: disable\n")
		file.WriteString("dbname: Goke_test\n")
		file.WriteString("sqlite_name: goke-test # SQLITE filename\n")
		file.WriteString("host: localhost\n")
		file.WriteString("port: 3306")
		if err != nil {
			panic("Failed to open file")
		}
	}

	defer file.Close()
}
func entryPointCmdRunner(cmd *cobra.Command, args []string) {

	figure := figure.NewFigure("Goke", "basic", true)
	figure.Print()
	fmt.Printf("Welcome, thanks for using Goke currently version is -v %s, to get started goke --help\n", VERSION)

	if initGoke {
		_, err := handlers.CreateMigrationsDir()
		if !os.IsNotExist(err) {

			fmt.Println("Started goke successfully")
		}

	}

}

func Execute() {

	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func addCommandFlags() {
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "Show version information")
	migrateCmd.Flags().StringVarP(&schema, "name", "n", utils.RandomMigrationName(), "Create history for migrations")
	dropCmd.Flags().StringVarP(&table, "table", "t", "", "Drop table if exists")
	dumpCmd.Flags().StringVarP(&tb_name, "json", "j", "goke_default", "Dump table schema to JSON file")
	RootCmd.Flags().BoolVarP(&initGoke, "init", "i", false, "Initialize goke")
}

func init() {
	addCommandFlags()

	RootCmd.AddCommand(migrateCmd)
	RootCmd.AddCommand(dumpCmd)
	RootCmd.AddCommand(dropCmd)
	RootCmd.AddCommand(testCmd)

}
