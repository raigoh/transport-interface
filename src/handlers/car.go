package handlers

import (
	"cars/src/data"
	"net/http"
	"path/filepath"
)

// CarHandler handles the HTTP request for the car details route.
// It extracts the car model ID from the URL, fetches the car model details,
// and renders the car details template. If any error occurs, it responds with
// the appropriate HTTP error status and message.
//
// Parameters:
//   - w (http.ResponseWriter): The response writer to send the HTTP response.
//   - r (*http.Request): The HTTP request.
//
// The URL should be in the format: /car/{id}
// Where {id} is the ID of the car model to retrieve.

func CarHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate an intentional error to test error handling
	// panic("intentional error for testing")

	id := r.URL.Path[len("/car/"):]
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	car, err := data.FetchModelByID(id)
	if err != nil {
		http.Error(w, "Failed to fetch car model: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Car data.CarModel
	}{
		Car: car,
	}

	RenderTemplate(w, filepath.Join("templates", "car.html"), data)
}
