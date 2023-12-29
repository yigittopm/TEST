package main

import (
	"log"
	"net/http"

	"github.com/yigittopm/test/database"
)

var (
	//cfg *config.Config
	err error
)

func init() {
	//cfg, err = config.LoadConfig()
	//if err != nil {
	//	panic(err)
	//}
}

func main() {
	// Init DB
	database.Start()

	http.HandleFunc("/user", database.GetUser1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
