package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Database DatabaseConfig
	}

	DatabaseConfig struct {
		Host     string
		Port     int
		Name     string
		User     string
		Password string
	}
)

func LoadConfig(env string) (Config, error) {

	viper.SetConfigFile(fmt.Sprintf("config/config_%s.yaml", env))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshalling config file: %v\n", err)
		return Config{}, err
	}

	return config, nil
}
