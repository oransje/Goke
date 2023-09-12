package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/handlers"
	"github.com/vsantos1/Goke/utils"
)

var tb_name string
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump table schema to json file",
	Run: func(cmd *cobra.Command, args []string) {

		c, err := config.ReadConfigFile("goke-config.yaml")
		utils.HandleError(err, "[Error]: reading file goke-config.yaml ")

		if tb_name == "goke_default" {
			log.Fatalln("[Error]: Please provide a table name to dump, use dump -j and provide table name")
		}

		typ := handlers.DumpSchema(tb_name, c)

		if typ.Problem != nil {
			log.Fatalf("[Error]: %v", typ.Problem)
		}

		println(typ.Message)
	},
}

func init() {

	dumpCmd.Flags().StringVarP(&tb_name, "json", "j", "goke_default", "Dump table schema to JSON file")
	RootCmd.AddCommand(dumpCmd)
}
