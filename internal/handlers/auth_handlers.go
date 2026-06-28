package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/krJay1/go-helpdesk/internal/models"
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
	token, err := h.getJWTtoken(&user)
	if err != nil {
		res.Error = err.Error()
		res.Send(w)
		return
	}
	res.Data = types.LoginResponse{User: &user, AccessToken: token}
	res.Status = http.StatusOK
	res.Message = "Login successfull."

	res.Send(w)
}

func (h *ApiHandler) getJWTtoken(u *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"email":   u.Email,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(h.JWTSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
