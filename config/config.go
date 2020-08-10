package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// AppConfig reflects the properties of the mysql database
type AppConfig struct {
	SQLDBConfig SQLDBConfig `yaml:"sqlDbConfig"`
}

// SQLDBConfig reflects the properties of the sql database
type SQLDBConfig struct {
	Type string `ỳaml:"type"`
	Host string `yaml:"host"`
	Port uint   `yaml:"port"`
}

// ReadConfig is responsible for read the config file
func ReadConfig(configFileName string) (*AppConfig, error) {
	var appConfig AppConfig
	file, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, &Error{Cause: err}
	}

	err = yaml.Unmarshal(file, &appConfig)
	if err != nil {
		return nil, &Error{Cause: err}
	}

	return &appConfig, nil
}

// Error is responsible for encapsulating errors generated by operations in the data access layer
type Error struct {
	Cause error
}

func (err *Error) Error() string {
	return err.Cause.Error()
}
