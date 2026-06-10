package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/krJay1/go-helpdesk/internal/handlers"
	"github.com/krJay1/go-helpdesk/internal/repository"
)

func InitializeUserRoutes(route *mux.Router, db *sql.DB) {

	userRepo := &repository.UserRepository{DB: db}

	userHandler := &handlers.UserHandler{Repo: userRepo}

	route.HandleFunc("/user", userHandler.CreateUserHandler).Methods("POST")
	route.HandleFunc("/user/{id}", userHandler.GetUserHandler).Methods("GET")

}
