package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnvVar(key string) string {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("invalid type assertion")
	}
	return value
}

type ConfigStruct struct {
	Username string
	Password string
	Smtp     string

	From string
	To   []string
}

var (
	configName  = ".env"
	configPaths = []string{".", "./config/"}
)

// Config is the package variable used to retrieve configuration options.
var Config = &ConfigStruct{}

// InitConfig initializes the configuration object with the
// variables from the configuration file or environment.
func InitConfig() {
	viper.SetConfigType("env")

	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}
	viper.SetConfigName(configName)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Reading configuration file failed: %v", err)
	}

	if err := viper.Unmarshal(Config); err != nil {
		log.Fatalf("Failed to unmarshall the configuration file: %v", err)
	}
}
