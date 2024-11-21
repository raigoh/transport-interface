package handlers

import (
	"cars/src/data"
	"net/http"
	"path/filepath"
)

// CompareHandler handles the HTTP request for the comparison route.
// It processes a POST request containing car model IDs to compare,
// fetches the details of each car model, and renders the comparison template.
// If any error occurs during the process, it responds with the appropriate HTTP error status and message.
//
// Parameters:
//   - w (http.ResponseWriter): The response writer to send the HTTP response.
//   - r (*http.Request): The HTTP request containing the car model IDs in the form data.
//
// The request should be a POST request with form data containing 'compare' parameters,
// where each parameter value is the ID of a car model to compare.
// The handler fetches details for each car model ID and renders a comparison page.

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate an intentional error to test error handling
	// panic("intentional error for testing")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	ids := r.Form["compare"]
	var cars []data.CarModel

	for _, id := range ids {
		car, err := data.FetchModelByID(id)
		if err != nil {
			http.Error(w, "Failed to fetch car model: "+err.Error(), http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}

	data := struct {
		Cars []data.CarModel
	}{
		Cars: cars,
	}

	RenderTemplate(w, filepath.Join("templates", "compare.html"), data)
}
