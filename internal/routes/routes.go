package routes

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/krJay1/go-helpdesk/internal/config"
	"github.com/krJay1/go-helpdesk/internal/handlers"
	"github.com/krJay1/go-helpdesk/internal/middleware"
	"github.com/krJay1/go-helpdesk/internal/repository"
)

func InitializeUserRoutes(route *mux.Router, db *pgxpool.Pool, cfg *config.Config) {

	appRepo := &repository.AppRepository{DB: db}

	apiHandler := handlers.NewApiHandler(appRepo, cfg)

	route.HandleFunc("/login", apiHandler.LoginHandler).Methods("POST")

	protectedRoute := route.PathPrefix("/").Subrouter()
	protectedRoute.Use(middleware.AuthMiddleWare)

	protectedRoute.HandleFunc("/user", apiHandler.CreateUserHandler).Methods("POST")
	protectedRoute.HandleFunc("/user/{id}", apiHandler.GetUserHandler).Methods("GET")
	protectedRoute.HandleFunc("/users", apiHandler.GetAllUsersHandler).Methods("GET")

}
