package handlers

import (
	"cars/src/data"
	"net/http"
	"path/filepath"
)

// Handler for the comparison route
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
