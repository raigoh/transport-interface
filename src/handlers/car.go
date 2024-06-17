package handlers

import (
	"cars/src/data"
	"net/http"
	"path/filepath"
)

// Handler for the car details route
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
