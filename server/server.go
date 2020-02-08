package server

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Server is a struct responsible for managing a *httprouter.Router.
// All HandlerFunc Closures hang off of this struct, so all HandlerFunc's
// have access to the server values.
type Server struct {
	Port   string
	Router *httprouter.Router
	DB     *sql.DB
}

// registerRoutes is responsible for wiring up all of our HandlerFunc
// to our server's router.
func (s *Server) registerRoutes() {
	s.Router.HandlerFunc(http.MethodGet, "/ping", s.ping())
}

// Start binds all routes to our router and then serves our
// router to handle all requests on incoming connections.
func (s *Server) Start() {
	s.registerRoutes()
	fmt.Println(fmt.Sprintf("Starting server on :%v", s.Port))
	http.ListenAndServe(":"+s.Port, s.Router)
}
