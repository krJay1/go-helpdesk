package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/krJay1/go-helpdesk/internal/models"
	"github.com/krJay1/go-helpdesk/internal/repository"
	"github.com/krJay1/go-helpdesk/internal/utils"
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
	response := utils.NewApiResponse()
	idstr := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := h.Repo.GetUser(id)
	if err != nil {
		response.Error = err
		response.Message = "Failed to fetch user"
		response.Send(w)
		return
	}

	response.Success = true
	response.Data = user
	response.Error = err
	response.Status = http.StatusOK
	response.Message = "User fetched successfully"

	response.Send(w)

}

func (h *UserHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	response := utils.NewApiResponse()
	response.Send(w)
}
