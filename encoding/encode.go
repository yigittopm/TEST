package a

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
}

func unmarshall(filename string) (user User, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File didn't open: %v\n", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("File didn't read: %v\n", err)
		return
	}
	json.Unmarshal(data, &user)
	return
}

func decoder(filename string) (user User, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("File didn't open: %v\n", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&user)
	if err != nil {
		fmt.Printf("File didn't read: %v\n", err)
		return
	}
	return
}
