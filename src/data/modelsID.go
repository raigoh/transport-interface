package data

import (
	"encoding/json"
	"net/http"
)

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
