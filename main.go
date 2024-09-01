package main

import (
	"cars/src/handlers"
	"cars/src/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Serve static files from the ./static directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve image files from the ./api/img directory under the /api/img/ route
	imgFS := http.FileServer(http.Dir("./api/img"))
	http.Handle("/api/img/", http.StripPrefix("/api/img/", imgFS))

	// Set up route handlers with recovery middleware
	http.HandleFunc("/compare", handlers.WithRecovery(handlers.CompareHandler))
	http.HandleFunc("/", handlers.WithRecovery(handlers.IndexHandler))
	http.HandleFunc("/car/", handlers.WithRecovery(handlers.CarHandler))
	http.HandleFunc("/health", handlers.WithRecovery(handlers.HealthCheckHandler))

	// Apply logging and rate limiting middleware to all routes
	loggedHandler := middleware.LoggingMiddleware(
		middleware.RateLimitingMiddleware(http.DefaultServeMux),
	)

	// Get the port from environment variable or use default 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a new http.Server with the configured handler
	server := &http.Server{
		Addr:    ":" + port,
		Handler: loggedHandler,
	}

	// Set up a channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		fmt.Println("Server is running on http://localhost:" + port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", port, err)
		}
	}()

	// Wait for interrupt signal
	<-stop

	// Create a context with a 5-second timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to shut down the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}

	log.Println("Server gracefully stopped")
}
