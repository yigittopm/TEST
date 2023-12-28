package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Start() {
	var err error

	source := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		"localhost",
		"postgres",
		"password",
		"godb",
		"5432",
		"disable",
		"Turkey",
	)

	DB, err = sql.Open("postgres", source)
	if err != nil {
		log.Fatalf("Db not connected: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Database succussfully connected.")
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func GetUser1(w http.ResponseWriter, r *http.Request) {
	var users []User
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Username)
		users = append(users, user)
	}

	usersBtye, _ := json.MarshalIndent(users, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(usersBtye)
}
