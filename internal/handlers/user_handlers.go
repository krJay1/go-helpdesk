package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/krJay1/go-helpdesk/internal/models"
	"github.com/krJay1/go-helpdesk/internal/repository"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := h.Repo.CreateUser(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)

}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idstr := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := h.Repo.GetUser(id)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)

}
