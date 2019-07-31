package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server ServerConfig
}

func InitServerConfig() ServerConfig {
	viper.SetConfigName("config.example")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var configuration Configuration
	err = viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return configuration.Server
}
