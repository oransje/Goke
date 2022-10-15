package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigYaml struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dialect  string `yaml:"dialect"`
	SslMode  string `yaml:"sslmode"`
	DbName   string `yaml:"dbname"`
	SqliteName string `yaml:"sqlite_name"`
	Host     string
	Port     int64
}

func ReadConfigFile(filename string) (*ConfigYaml, error) {

	buf, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("error, while reading file %v", filename)
	}

	config := &ConfigYaml{}

	err = yaml.Unmarshal(buf, config)

	if err != nil {
		return nil, fmt.Errorf("failed to deserialize %s.yaml", filename)
	}

	return config, err
}
