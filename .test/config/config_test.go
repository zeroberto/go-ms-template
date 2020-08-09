package config

import (
	"testing"

	"github.com/zeroberto/go-ms-template/config"
)

func TestReadConfig(t *testing.T) {
	expectedAppConfig := config.AppConfig{
		SQLDBConfig: config.SQLDBConfig{
			Type: "test",
			Host: "host",
			Port: 1,
		},
	}

	configFileName := "applicationTest.yml"
	appConfig, err := config.ReadConfig(configFileName)

	if err != nil {
		t.Errorf("ReadConfig() failed, error %v", err)
	}

	if expectedAppConfig != *appConfig {
		t.Errorf("ReadConfig() failed, expected %v, got %v", expectedAppConfig, appConfig)
	}
}
