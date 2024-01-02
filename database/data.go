package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/yigittopm/test/config"
)

func Start(cfg config.Config) (*sql.DB, error) {
	var err error

	source := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		cfg.DB_HOST,
		cfg.DB_USER,
		cfg.DB_PASSWORD,
		cfg.DB_NAME,
		cfg.DB_PORT,
		cfg.DB_SSLMODE,
		cfg.DB_TIMEZONE,
	)
	fmt.Println(source)
	//source := fmt.Sprintf(
	//	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
	//	"psql",
	//	"postgres",
	//	"password",
	//	"godb",
	//	"5432",
	//	"disable",
	//	"Turkey",
	//)

	DB, err := sql.Open("postgres", source)
	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, nil
}
