package main

import (
	"github.com/vsantos1/Goke/cmd"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/database"
)

func main() {

	cmd.RootCmd.Use = "goke"
	cmd.Execute()
	c, err := config.ReadConfigFile("goke-config.yaml")
	database.ConnectionDatabase(c)
	if err != nil {
		panic("Error, file format is invalid delete the old goke-config.yaml")
	}
}
