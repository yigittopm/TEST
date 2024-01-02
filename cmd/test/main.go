package main

import (
	"context"
	"os"

	"github.com/yigittopm/test/config"
	"github.com/yigittopm/test/internal/app"
)

func main() {
	env := os.Getenv("env")
	if env == "" {
		env = "dev"
	}

	cfg, err := config.LoadConfig(env)
	if err != nil {
		panic(err)
	}

	app := app.NewApp(context.Background(), cfg)
	if err := app.Start(); err != nil {
		panic(err)
	}
}
