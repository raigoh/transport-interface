package data

import (
	"encoding/json"
	"net/http"
)

// FetchModelByID retrieves a car model by its ID from the API.
// It sends a GET request to the /api/models/{id} endpoint and decodes the JSON response
// into a CarModel struct. Additionally, it fetches the manufacturer and category details
// for the car model.
//
// Parameters:
//   - id (string): The ID of the car model to retrieve.
//
// Returns:
//   - (CarModel, error): A CarModel struct and an error.
//     If the request or decoding fails, the error will be non-nil.

func FetchModelByID(id string) (CarModel, error) {
	response, err := http.Get("http://localhost:3000/api/models/" + id)
	if err != nil {
		return CarModel{}, err
	}
	defer response.Body.Close()

	var model CarModel
	if err := json.NewDecoder(response.Body).Decode(&model); err != nil {
		return CarModel{}, err
	}

	// Fetch manufacturer details
	manufacturer, err := FetchManufacturerByID(model.ManufacturerID)
	if err != nil {
		return CarModel{}, err
	}
	model.ManufacturerDetails = manufacturer

	// Fetch category details
	category, err := FetchCategoryByID(model.CategoryID)
	if err != nil {
		return CarModel{}, err
	}
	model.CategoryDetails = category

	return model, nil
}
