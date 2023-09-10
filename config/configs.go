package config

import (
	"os"

	"github.com/vsantos1/Goke/utils"
	"gopkg.in/yaml.v2"
)

type ConfigYaml struct {
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Dialect    string `yaml:"dialect"`
	SslMode    string `yaml:"sslmode"`
	DbName     string `yaml:"dbname"`
	SqliteName string `yaml:"sqlite_name"`
	Host       string `yaml:"host"`
	Port       int64  `yaml:"port"`
}

func ReadConfigFile(filename string) (*ConfigYaml, error) {

	buf, err := os.ReadFile(filename)

	utils.HandleError(err, "[ERROR]: reading goke-config.yaml file")

	config := &ConfigYaml{}

	err = yaml.Unmarshal(buf, config)

	utils.HandleError(err, "[ERROR]: unmarshalling goke-config.yaml file")

	return config, err
}
