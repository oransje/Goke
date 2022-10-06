package main

import (
	"github.com/vsantos1/Goke/cmd"
	"github.com/vsantos1/Goke/config"
)

func main() {
	cmd.RootCmd.Use = "goke"
	cmd.Execute()

	_, err := config.ReadConfigFile("goke-config.yaml")
	if err != nil {
		panic("Error, file format is invalid delete the old goke-config.yaml")
	}
}
