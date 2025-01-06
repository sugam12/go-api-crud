package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	user "github.com/sugam12/go-api-crud/service"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	userService := user.NewHandler()
	userService.RegisterRoutes(subrouter)
	log.Println("listening in port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
