package handlers

import (
	"context"
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

	err = utils.VerifyPassword(req.Password, user.PasswordHash)
	if err != nil {
		res.Error = "Password verification failed"
		res.Send(w)
		return
	}

	accessToken, err := h.getAccessToken(&user)
	if err != nil {
		res.Error = err.Error()
		res.Send(w)
		return
	}

	refreshToken, err := h.getRefereshToken(ctx, user.ID)
	if err != nil {
		res.Error = err.Error()
		res.Send(w)
		return
	}

	res.Data = types.LoginResponse{User: &user, AccessToken: accessToken, RefreshToken: refreshToken}
	res.Status = http.StatusOK
	res.Message = "Login successfull."

	res.Send(w)
}

func (h *ApiHandler) getAccessToken(u *models.User) (string, error) {
	expTime := time.Now().Add(15 * time.Minute)
	claims := jwt.MapClaims{
		"user_id": u.ID,
		"email":   u.Email,
		"exp":     expTime.Unix(),
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(h.JWTSecret)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (h *ApiHandler) getRefereshToken(ctx context.Context, userId int64) (string, error) {
	expTime := time.Now().Add(time.Hour)
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     expTime.Unix(),
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(h.JWTSecret)
	if err != nil {
		return "", err
	}
	err = h.Repo.UpsertRefreshToken(ctx, userId, token, expTime)
	if err != nil {
		return "", err
	}
	return token, nil
}
