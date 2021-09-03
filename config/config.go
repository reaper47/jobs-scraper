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
