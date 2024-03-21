package usecase

import (
	"database/sql"
	"net/http"
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

	tests := []struct {
		description         string
		route               string
		method              string
		expectedCode        int
		expectedContentType string
		expectedBody        any
		body                any
	}{
		{
			description:         "New user registration with valid data",
			route:               "/api/v1/users/profile",
			method:              "GET",
			expectedCode:        404,
			expectedContentType: "application/json",
			expectedBody:        nil,
			body:                nil,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(
			test.method,
			test.route,
			nil,
		)

		res, err := app.Test(req, -1)
		if err != nil {
			t.Errorf("ERROR: %#v", err)
		}

		if res.StatusCode != test.expectedCode {
			t.Errorf("Response Code: %#v != Expected Code: %#v\n", res.StatusCode, test.expectedCode)
		}

		contentType := res.Header.Get("Content-Type")
		if contentType != test.expectedContentType && res.StatusCode == 200 {
			t.Errorf("Content-Type: %#v != %#v", contentType, test.expectedContentType)
		}

	}

}

func TestLogin(t *testing.T) {

}

func TestLogout(t *testing.T) {

}

func TestProfile(t *testing.T) {

}
