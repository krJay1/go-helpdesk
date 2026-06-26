package routes

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krJay1/go-helpdesk/internal/handlers"
	"github.com/krJay1/go-helpdesk/internal/repository"
)

func InitializeUserRoutes(route *mux.Router, db *pgxpool.Pool) {

	appRepo := &repository.AppRepository{DB: db}

	apiHandler := handlers.NewUserHandler(appRepo)

	route.HandleFunc("/login", apiHandler.LoginHandler).Methods("POST")

	route.HandleFunc("/user", apiHandler.CreateUserHandler).Methods("POST")
	route.HandleFunc("/user/{id}", apiHandler.GetUserHandler).Methods("GET")
	route.HandleFunc("/users", apiHandler.GetAllUsersHandler).Methods("GET")

}
