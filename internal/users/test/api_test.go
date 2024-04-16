package usecase

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/yigittopm/wl-auth/internal/users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	app    *fiber.App
	mockDb *sql.DB
	mock   sqlmock.Sqlmock
)

func Setup(t *testing.T) {
	t.Helper()
	app = fiber.New()

	mockDb, mock, _ = sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})

	//repo := usersRepository.New(db)
	//usecase := usersUsecase.New(repo)
	//handler = usersHandler.New(usecase)

	version := app.Group("/api/v1")
	users.Setup(version, db)
}

func TestRegister(t *testing.T) {
	Setup(t)
}

func TestLogin(t *testing.T) {

}

func TestLogout(t *testing.T) {

}

func TestProfile(t *testing.T) {

}
