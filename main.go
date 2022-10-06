package main

import (
	"fmt"

	"github.com/vsantos1/Goke/cmd"
	"github.com/vsantos1/Goke/config"
)

func main() {
	cmd.Init()
	c, err := config.ReadConfigFile("config.yaml")
	if err != nil {
		panic("Error, file format is invalid delete the old config.yaml")
	}

	fmt.Printf("%#v", c)

}
