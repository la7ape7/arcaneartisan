package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Environment EnvironmentConfiguration
	Database    DatabaseConfiguration
}

type EnvironmentConfiguration struct {
	Type         string
	Port         string
	ApiSecretKey string
}

type DatabaseConfiguration struct {
	SqliteFilePath string
}

func GetDefaultAppConfig() *Configuration {
	var defaultConfiguration Configuration
	defaultConfiguration.Environment.Port = "8080"
	defaultConfiguration.Environment.Type = "prod"
	defaultConfiguration.Environment.ApiSecretKey = "SecretKey"
	defaultConfiguration.Database.SqliteFilePath = "wiw.sqlite3"
	return &defaultConfiguration
}

func GetConfiguration() *Configuration {
	var appConfig Configuration
	if _, err := toml.DecodeFile("./config/config.toml", &appConfig); err != nil {
		fmt.Println("Can't read configuration file, will use default configuration")
		return GetDefaultAppConfig()
	}
	return &appConfig
}
