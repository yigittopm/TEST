package database

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() (*gorm.DB, error) {
	_, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		os.Setenv("POSTGRES_HOST", "localhost")
		os.Setenv("POSTGRES_USER", "postgres")
		os.Setenv("POSTGRES_PASSWORD", "password")
		os.Setenv("POSTGRES_DB", "godb")
		os.Setenv("POSTGRES_PORT", "5432")
		os.Setenv("POSTGRES_SSLMODE", "disable")
	}

	source := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
	)

	DB, err := gorm.Open(postgres.Open(source), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
