package routes

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krJay1/go-helpdesk/internal/handlers"
	"github.com/krJay1/go-helpdesk/internal/repository"
)

func InitializeUserRoutes(route *mux.Router, db *pgxpool.Pool) {

	userRepo := &repository.UserRepository{DB: db}

	userHandler := &handlers.UserHandler{Repo: userRepo}

	route.HandleFunc("/user", userHandler.CreateUserHandler).Methods("POST")
	route.HandleFunc("/user/{id}", userHandler.GetUserHandler).Methods("GET")
	route.HandleFunc("/users", userHandler.GetAllUsersHandler).Methods("GET")

}
