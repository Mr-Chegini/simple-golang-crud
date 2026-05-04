package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Mr-Chegini/simple-golang-crud/database"
)

// Response represents a JSON response
type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message"`
	Version string `json:"version,omitempty"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Status:  "ok",
		Message: "Simple Go CRUD API is running",
	}
	json.NewEncoder(w).Encode(response)
}

func apiRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Welcome to Simple Go CRUD API",
		Version: "v1",
	}
	json.NewEncoder(w).Encode(response)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	// Run database migrations
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", healthHandler)

	// API routes
	mux.HandleFunc("/api/v1/", apiRootHandler)

	// Apply middleware
	handler := corsMiddleware(loggingMiddleware(mux))

	fmt.Println("🚀 Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
