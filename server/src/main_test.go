package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AppHandlerFake struct{}

func (a AppHandlerFake) SetupRoutes(app *fiber.App) {
	setupRoutes(app)
}

var mock sqlmock.Sqlmock

func (a *AppHandlerFake) InitDb() {
	mockDb, _mock, _ := sqlmock.New()
	mock = _mock

	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	database.DB = database.Dbinstance{
		Db: db,
	}
}

func TestApp(t *testing.T) {

	testTable := []struct {
		name        string
		method      string
		path        string
		statusCode  int
		requestBody map[string]interface{}
		expectFunc  func()
	}{
		{
			name:        `GET /favorites`,
			method:      http.MethodGet,
			path:        `/favorites`,
			statusCode:  200,
			requestBody: map[string]interface{}{},
			expectFunc:  func() { mock.ExpectQuery(`SELECT`).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "image"})) },
		},
		{
			name:        `GET /favorites/id`,
			method:      http.MethodGet,
			path:        `/favorites/1`,
			statusCode:  200,
			requestBody: map[string]interface{}{},
			expectFunc: func() {
				mock.ExpectQuery(`SELECT`).WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"id", "title", "image"}))
			},
		},
	}

	app := CreateApp(&AppHandlerFake{})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {

			rbody, _ := json.Marshal(tc.requestBody)
			request := httptest.NewRequest(tc.method, tc.path, bytes.NewReader(rbody))
			request.Header.Add(`Content-Type`, `application/json`)

			tc.expectFunc()

			response, _ := app.Test(request)

			statusCode := response.StatusCode
			if statusCode != tc.statusCode {
				t.Errorf("StatusCode was incorrect, got: %d, want: %d.", statusCode, tc.statusCode)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}

		})
	}
}
