package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/krJay1/go-helpdesk/internal/types"
	"github.com/krJay1/go-helpdesk/internal/utils"
)

func (h *ApiHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res := utils.NewApiResponse()
	var req types.LoginPayload

	json.NewDecoder(r.Body).Decode(&req)

	user, err := h.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		res.Error = err.Error()
		res.Send(w)
		return
	}

	passwordVerified := utils.VerifyPassword(req.Password, user.PasswordHash)
	if !passwordVerified {
		res.Error = "Password verification failed"
		res.Send(w)
		return
	}
	res.Data = user

	res.Send(w)
}
