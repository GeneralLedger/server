package server

import (
	"net/http"

	"github.com/generalledger/response"
)

func (s *Server) ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := response.New(w)
		defer resp.Output()
		resp.SetResult(http.StatusOK, "pong")
	}
}
