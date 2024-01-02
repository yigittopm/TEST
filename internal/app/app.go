package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/yigittopm/test/config"
	"github.com/yigittopm/test/database"

	userV1 "github.com/yigittopm/test/internal/users/delivery/http/v1"
	userRepository "github.com/yigittopm/test/internal/users/repository"
	userUsecase "github.com/yigittopm/test/internal/users/usecase"
)

type App struct {
	DB   *sql.DB
	Echo *echo.Echo
	Cfg  config.Config
}

func NewApp(ctx context.Context, cfg config.Config) *App {
	db, err := database.Start(cfg)
	if err != nil {
		panic(err)
	}

	return &App{
		DB:   db,
		Echo: echo.New(),
		Cfg:  cfg,
	}
}

func (app *App) Start() error {
	if err := app.StartService(); err != nil {
		return err
	}

	//app.Echo.Debug = app.Cfg.Server.Debug
	//app.Echo.Use(middleware.AppCors())
	//app.Echo.Use(middleware.CacheWithRevalidation)
	return app.Echo.StartServer(&http.Server{
		Addr:         fmt.Sprintf(":8080"),
		ReadTimeout:  3000,
		WriteTimeout: 3000,
	})
}

func (app *App) StartService() error {
	userRepo := userRepository.New(app.DB)

	userUsecase := userUsecase.New(userRepo, app.Cfg)

	userController := userV1.New(userUsecase)

	version := app.Echo.Group("/api/v1/")

	userV1.UserRoute(version, userController, app.Cfg)

	return nil
}
