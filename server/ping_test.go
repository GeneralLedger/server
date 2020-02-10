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
	// Prepare Test Database
	db, err := database.MigratedDatabase(os.Getenv("TEST_POSTGRES_URL"))
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
			Result:     "pong",
		},
		response.Parse(recorder.Result().Body),
	)
}
