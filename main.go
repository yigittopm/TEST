package main

import (
	"github.com/yigittopm/test/api"
)

func main() {
	a := api.SetupServer()
	a.Run(":8080")
}
