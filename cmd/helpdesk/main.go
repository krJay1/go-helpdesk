package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/krJay1/go-helpdesk/internal/storage"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func main() {
	fmt.Println("Hello World")

	connectionStr := "host=localhost dbname=go-db connect_timeout=5"
	err := storage.InitDB(connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.DB.Close()

	root := mux.NewRouter()
	root.HandleFunc("/", HomeHandler)

	addr := ":8088"

	log.Println("Server is starting on port :8088")
	server := http.Server{
		Addr:    addr,
		Handler: root,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

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
