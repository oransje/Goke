package cmd

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/handlers"
)

var version bool
var initGoke bool
var RootCmd = &cobra.Command{
	SilenceUsage: true,
	Short:        "Goke is schema migrations for golang's database/sql",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Create base configuration to connect on database if not exists
		file, err := os.OpenFile("goke-config.yaml", 0, 0644)
		if os.IsNotExist(err) {
			file, err := os.Create("goke-config.yaml")
			file.WriteString("datasource: root:root@(localhost:3306)/Goke_test?charset=utf8mb4&parseTime=true\n")
			file.WriteString("username: root\n")
			file.WriteString("password: root\n")
			file.WriteString("dialect: mysql\n")
			file.WriteString("host: localhost\n")
			file.WriteString("port: 3306")
			if err != nil {
				panic("Failed to open file")
			}
		}

		defer file.Close()
	},
	Run: func(cmd *cobra.Command, args []string) {

		figure := figure.NewFigure("Goke", "basic", true)
		figure.Print()
		fmt.Printf("Welcome, thanks for using Goke currently version is -v %s, to get started goke --help\n", VERSION)

		if initGoke {
			_, err := handlers.CreateMigrationsDir()
			if !os.IsNotExist(err) {

				fmt.Println("Started goke successfully")
			}

		}

	},
}

func Execute() {

	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
func init() {
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "Show version information")
	RootCmd.Flags().BoolVarP(&initGoke, "init", "i", false, "Init goke main structure")

}
