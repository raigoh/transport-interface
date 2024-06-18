package data

import (
	"encoding/json"
	"net/http"
)

// FetchModels retrieves the list of car models from the API.
// It sends a GET request to the /api/models endpoint and decodes the JSON response
// into a slice of CarModel structs.
//
// Returns:
//   - ([]CarModel, error): A slice of CarModel structs and an error.
//     If the request or decoding fails, the error will be non-nil.

func FetchModels() ([]CarModel, error) {
	response, err := http.Get("http://localhost:3000/api/models")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var models []CarModel
	if err := json.NewDecoder(response.Body).Decode(&models); err != nil {
		return nil, err
	}
	return models, nil
}
