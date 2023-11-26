package main

import "github.com/yigittopm/test/api"

func main() {
	a := api.Server{}
	server := a.NewServer(":8080")
	server.Start()
}
