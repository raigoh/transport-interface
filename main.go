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
	// Serves static files from the local directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serves image files from ./api/img under the /api/img/ route.
	imgFS := http.FileServer(http.Dir("./api/img"))
	http.Handle("/api/img/", http.StripPrefix("/api/img/", imgFS))

	http.HandleFunc("/compare", handlers.WithRecovery(handlers.CompareHandler))
	http.HandleFunc("/", handlers.WithRecovery(handlers.IndexHandler))
	http.HandleFunc("/car/", handlers.WithRecovery(handlers.CarHandler))
	http.HandleFunc("/health", handlers.WithRecovery(handlers.HealthCheckHandler))

	loggedHandler := middleware.LoggingMiddleware(
		middleware.RateLimitingMiddleware(http.DefaultServeMux),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: loggedHandler,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("Server is running on http://localhost:" + port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", port, err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v\n", err)
	}

	log.Println("Server gracefully stopped")
}
