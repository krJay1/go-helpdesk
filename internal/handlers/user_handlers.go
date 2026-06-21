package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/krJay1/go-helpdesk/internal/models"
	"github.com/krJay1/go-helpdesk/internal/repository"
	"github.com/krJay1/go-helpdesk/internal/types"
	"github.com/krJay1/go-helpdesk/internal/utils"
)

func NewUserHandler(repository *repository.AppRepository) *ApiHandler {
	return &ApiHandler{
		Repo: repository,
	}
}

func (h *ApiHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	res := utils.NewApiResponse()
	var payload types.CreateUser

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		res.Error = err.Error()
		res.Send(w)
		return
	}
	passwordHash, err := utils.HashPassword(payload.Password)
	if err != nil {
		res.Error = err.Error()
		res.Send(w)
		return
	}
	user := models.User{
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
		Email:        payload.Email,
		PasswordHash: passwordHash,
		MobileNumber: &payload.MobileNumber,
	}

	id, err := h.Repo.CreateUser(user)

	if err != nil {
		res.Error = err.Error()
		res.Message = "Failed to create users acccount"
		res.Send(w)
		return
	}

	user.ID = id

	res.Status = http.StatusOK
	res.Data = user
	res.Success = true
	res.Message = "User Created Successfully"

	res.Send(w)
}

func (h *ApiHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	response := utils.NewApiResponse()
	idstr := mux.Vars(r)["id"]

	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		response.Send(w)
		return
	}

	user, err := h.Repo.GetUserByID(id)
	if err != nil {
		response.Error = err.Error()
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

func (h *ApiHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	res := utils.NewApiResponse()
	users, err := h.Repo.GetUsers()
	if err != nil {
		res.Error = err.Error()
		res.Send(w)
		return
	}

	res.Success = true
	res.Data = users
	res.Message = "Records fetched successfully"
	res.Status = http.StatusOK

	res.Send(w)
}
