package product

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	types "github.com/sugam12/go-api-crud/payload"
	"github.com/sugam12/go-api-crud/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/product", h.HandleProduct)
}

func (h *Handler) HandleProduct(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.HandleGetProduct(w, r)
	case http.MethodPost:
		h.HandlePostProduct(w, r)
	case http.MethodPut:
		h.HandleUpdateProduct(w, r)
	default:
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("method not allowed"))
	}
}

func (h *Handler) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	prod, err := h.store.GetProduct()
	if err != nil {
		log.Fatal(err)
	}
	utils.WriteJSON(w, http.StatusOK, prod)

}
func (h *Handler) HandlePostProduct(w http.ResponseWriter, r *http.Request) {
	var payload types.Product
	err := utils.ParseJSON(r, payload)
	if err != nil {
		log.Fatal(err)
	}

	prod := new(types.Product)
	prod, err = h.store.CreateProduct(payload)
	if err != nil {
		log.Fatal(err)
	}
	utils.WriteJSON(w, http.StatusCreated, prod)

}

func (h *Handler) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {

}
