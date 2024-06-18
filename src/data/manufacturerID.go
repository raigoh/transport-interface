package data

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// FetchManufacturerByID retrieves a manufacturer by its ID from the API.
// It sends a GET request to the /api/manufacturers/{id} endpoint and decodes the JSON response
// into a Manufacturer struct.
//
// Parameters:
//   - id (int): The ID of the manufacturer to retrieve.
//
// Returns:
//   - (Manufacturer, error): A Manufacturer struct and an error.
//     If the request or decoding fails, the error will be non-nil.

func FetchManufacturerByID(id int) (Manufacturer, error) {
	response, err := http.Get("http://localhost:3000/api/manufacturers/" + strconv.Itoa(id))
	if err != nil {
		return Manufacturer{}, err
	}
	defer response.Body.Close()
	var manufacturer Manufacturer
	if err := json.NewDecoder(response.Body).Decode(&manufacturer); err != nil {
		return Manufacturer{}, err
	}
	return manufacturer, nil
}
