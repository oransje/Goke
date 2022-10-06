package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goke",
	Short: "Goke is schema migrations for golang's database/sql",
	Long: `A Fast and simple data migration for Golang built with love by vsantos1 and friends in Go.
		Complete documentation is available at https://github.com/vsantos1/Goke`,
	Run: func(cmd *cobra.Command, args []string) {

		// Create base configuration to connect on database if not exists

		file, err := os.OpenFile("config.yaml", 0, 0644)
		if os.IsNotExist(err) {

			file, err := os.Create("config.yaml")
			file.WriteString("datasource: root:root@(localhost:3306)database?charset=utf8mb4&parseTime=true\n")
			file.WriteString("username: root\n")
			file.WriteString("password: root\n")
			file.WriteString("dialect: mysql\n")
			file.WriteString("host: localhost\n")
			file.WriteString("port: 3306")
			if err != nil {
				panic(err)
			}
		}

		defer file.Close()

	},
}

func Init() {
	// rootCmd.Flags().StringVarP(&FileName, "init", "i", "config", "Database configuration file, should be FILENAME.yaml")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
