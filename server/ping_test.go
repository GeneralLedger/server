package server

import (
	"github.com/generalledger/api/database"
	"github.com/generalledger/response"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestPingSuccess(t *testing.T) {
	// Get a connection to the database
	db, err := database.NewConnection(
		database.ConnectionInfo{
			Host:         os.Getenv("TEST_POSTGRES_HOST"),
			Port:         os.Getenv("TEST_POSTGRES_PORT"),
			User:         os.Getenv("TEST_POSTGRES_USER"),
			Password:     os.Getenv("TEST_POSTGRES_PASSWORD"),
			DatabaseName: os.Getenv("TEST_POSTGRES_DB"),
			SSLMode:      os.Getenv("TEST_POSTGRES_SSL_MODE"),
		}.ToURI(),
		database.Config{},
	)
	if err != nil {
		panic(err)
	}

	// Prepare Server
	server := &Server{
		DB: db,
	}

	// Send Request
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	server.ping()(recorder, req)

	// Test
	assert.Equal(t,
		response.Response{
			StatusCode: 200,
			StatusText: "OK",
			Result: map[string]interface{}{
				"database_connection": "OK",
			},
		},
		response.Parse(recorder.Result().Body),
	)
}
