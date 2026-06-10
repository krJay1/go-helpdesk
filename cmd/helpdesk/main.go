package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/krJay1/go-helpdesk/internal/config"
	"github.com/krJay1/go-helpdesk/internal/handlers"
	"github.com/krJay1/go-helpdesk/internal/routes"
	"github.com/krJay1/go-helpdesk/internal/storage"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func main() {
	cfg := config.Load()

	db, err := storage.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	root := mux.NewRouter()
	root.HandleFunc("/", HomeHandler)
	root.HandleFunc("/health", handlers.HealthCheck).Methods("GET")

	routes.InitializeUserRoutes(root, db)

	addr := ":8088"

	log.Println("Server is starting...")
	server := http.Server{
		Addr:    addr,
		Handler: root,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	log.Println("✅ Server is running on http://localhost:8088")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	log.Println("Shutting down server...")

	// Give active request 5 seconds to execute
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}

	log.Println("Server shutdown gracefully.")
}
