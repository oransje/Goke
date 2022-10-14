package cmd

import (
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/handlers"
)

var tb_name string
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump table schema to json file",
	Run: func(cmd *cobra.Command, args []string) {
		time.Sleep(1000)

		c, err := config.ReadConfigFile("goke-config.yaml")
		if err != nil {
			panic("Error while reading file goke-config..yaml ")
		}
		if tb_name == "goke_default" {
			log.Fatalln("Please provide a table name to dump, use --dump or -d provide a table name")
		}
		typ := handlers.DumpSchema(tb_name,c)

		if typ.Problem != nil {
			panic("Error while reading file goke-config..yaml ")
		}

		println(typ.Message)
	},
}

func init() {

	dumpCmd.Flags().StringVarP(&tb_name, "json", "j", "goke_default", "Dump table schema to json") 
	RootCmd.AddCommand(dumpCmd)
}
