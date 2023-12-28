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
		DB_HOST     string `mapstructure:"POSTGRES_HOST"`
		DB_USER     string `mapstructure:"POSTGRES_USER"`
		DB_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
		DB_NAME     string `mapstructure:"POSTGRES_DB"`
		DB_PORT     string `mapstructure:"POSTGRES_PORT"`
		DB_SSLMODE  string `mapstructure:"POSTGRES_SSLMODE"`
		DB_TIMEZONE string `mapstructure:"POSTGRES_TIMEZONE"`
	}
)

func LoadConfig() (Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

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
