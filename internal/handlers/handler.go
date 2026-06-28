package handlers

import (
	"github.com/krJay1/go-helpdesk/internal/config"
	"github.com/krJay1/go-helpdesk/internal/repository"
)

type ApiHandler struct {
	Repo *repository.AppRepository
	*config.Config
}

func NewApiHandler(repository *repository.AppRepository, cfg *config.Config) *ApiHandler {
	return &ApiHandler{
		Repo:   repository,
		Config: cfg,
	}
}
