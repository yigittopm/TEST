package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Start() (*sql.DB, error) {
	//source := fmt.Sprintf(
	//	"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
	//	os.Getenv("POSTGRES_HOST"),
	//	os.Getenv("POSTGRES_USER"),
	//	os.Getenv("POSTGRES_PASSWORD"),
	//	os.Getenv("POSTGRES_DB"),
	//	os.Getenv("POSTGRES_PORT"),
	//	os.Getenv("SSLMODE"),
	//	os.Getenv("POSTGRES_TIMEZONE"),
	//)

	DB, err := sql.Open("postgres", "postgres://postgres:password@localhost/godb?sslmode=disable")
	if err != nil {
		return nil, err
	}

	//AutoMigrate(DB)

	return DB, nil
}

func AutoMigrate(DB *sql.DB) {
	_, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id serial primary key,
		username varchar(255) not null,
		email varchar(255) not null,
		password varchar(255) not null,
		user_type varchar(255),
		is_active boolean,
		created_at TIMESTAMP not null default NOW(),
		created_by varchar(255) not null,
		updated_at TIMESTAMP not null default NOW(),
		updated_by varchar(255) not null
	);
	`)
	fmt.Print(err)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success created users table")
}
