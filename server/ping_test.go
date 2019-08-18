package server

import (
	"github.com/generalledger/response"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingSuccess(t *testing.T) {
	// Prepare Server
	server := &Server{}

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
