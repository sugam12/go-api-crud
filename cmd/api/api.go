package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sugam12/go-api-crud/service/product"
	"github.com/sugam12/go-api-crud/service/user"
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

	userStore := user.NewStore(s.db)
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productService := product.NewHandler(productStore)
	productService.RegisterRoutes(subrouter)

	log.Println("listening in port", s.addr)
	return http.ListenAndServe(s.addr, router)
}
