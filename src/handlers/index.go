package handlers

import (
	"cars/src/data"
	"cars/src/filters"
	"net/http"
	"path/filepath"
)

// Handler for the index route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate an intentional error to test error handling
	// panic("intentional error for testing")

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		manufacturer := r.Form.Get("manufacturer")
		category := r.Form.Get("category")
		minYear := r.Form.Get("minYearSelect")
		maxYear := r.Form.Get("maxYearSelect")
		searchQuery := r.Form.Get("searchQuery")

		models, err := filters.FetchModelsFiltered(manufacturer, minYear, maxYear, category, searchQuery)
		if err != nil {
			http.Error(w, "Failed to fetch models", http.StatusInternalServerError)
			return
		}

		categories, err := data.FetchCategories()
		if err != nil {
			http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
			return
		}
		manufacturers, err := data.FetchManufacturers()
		if err != nil {
			http.Error(w, "Failed to fetch manufacturers", http.StatusInternalServerError)
			return
		}
		filterData, err := filters.FetchFilterData()
		if err != nil {
			http.Error(w, "Failed to fetch filter data", http.StatusInternalServerError)
			return
		}

		data := struct {
			Models        []data.CarModel
			Categories    []data.Category
			Manufacturers []data.Manufacturer
			FilterData    data.FilterData
		}{
			Models:        models,
			Categories:    categories,
			Manufacturers: manufacturers,
			FilterData:    filterData,
		}

		RenderTemplate(w, filepath.Join("templates", "index.html"), data)
	} else if r.Method == http.MethodGet {
		models, err := data.FetchModels()
		if err != nil {
			http.Error(w, "Failed to fetch models", http.StatusInternalServerError)
			return
		}

		categories, err := data.FetchCategories()
		if err != nil {
			http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
			return
		}
		manufacturers, err := data.FetchManufacturers()
		if err != nil {
			http.Error(w, "Failed to fetch manufacturers", http.StatusInternalServerError)
			return
		}
		filterData, err := filters.FetchFilterData()
		if err != nil {
			http.Error(w, "Failed to fetch filter data", http.StatusInternalServerError)
			return
		}

		data := struct {
			Models        []data.CarModel
			Categories    []data.Category
			Manufacturers []data.Manufacturer
			FilterData    data.FilterData
		}{
			Models:        models,
			Categories:    categories,
			Manufacturers: manufacturers,
			FilterData:    filterData,
		}

		RenderTemplate(w, filepath.Join("templates", "index.html"), data)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
