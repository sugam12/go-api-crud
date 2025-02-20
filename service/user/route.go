package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sugam12/go-api-crud/config"
	types "github.com/sugam12/go-api-crud/payload"
	"github.com/sugam12/go-api-crud/service/auth"
	"github.com/sugam12/go-api-crud/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.HandleLogin).Methods("POST")
	router.HandleFunc("/register", h.HandleRegister).Methods("POST")
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginPayload
	err := utils.ParseJSON(r, payload)
	if err != nil {
		log.Fatal(err)
	}

	user := new(types.User)
	user, err = h.store.GetUserByEmail(payload.UserName)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s does not exists", payload.UserName))
		return
	}
	error := auth.ComparePassword(user.Password, payload.Password)
	if error != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("password of email %s do not match", payload.UserName))
		return
	}
	secret := []byte(config.EnvVars.JWTSecret)
	token, err := auth.CreateJWT(secret, user.Id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

	//create a jwt token for 24 hours and send it back as response
}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	var payload types.RegisterPayload

	err := utils.ParseJSON(r, payload)
	if err != nil {
		log.Fatal(err)
	}

	// check if user exists by email
	_, err = h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	// password hash
	hashedPassword, err := auth.HashedPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	newUser := new(types.User)
	// create new user
	newUser, err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, newUser)

}
