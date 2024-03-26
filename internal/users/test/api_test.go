package usecase

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/yigittopm/wl-auth/internal/users"
	"github.com/yigittopm/wl-auth/internal/users/dtos"
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
			route:               "/users/register",
			method:              "POST",
			expectedCode:        200,
			expectedContentType: "application/json",
			expectedBody:        nil,
			body: dtos.RegisterRequest{
				Email:    "yigittopm@hotmail.com",
				Username: "yigittopm",
				Password: "password",
			},
		},
	}

	for _, test := range tests {
		jsonData, _ := json.Marshal(test.body)
		req, _ := http.NewRequest(
			test.method,
			test.route,
			bytes.NewBuffer(jsonData),
		)
		req.Header.Set("Content-Type", "application/json")

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
