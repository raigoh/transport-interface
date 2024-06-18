package data

import (
	"encoding/json"
	"net/http"
)

// FetchManufacturers retrieves the list of manufacturers from the API.
// It sends a GET request to the /api/manufacturers endpoint and decodes the JSON response
// into a slice of Manufacturer structs.
//
// Returns:
//   - ([]Manufacturer, error): A slice of Manufacturer structs and an error.
//     If the request or decoding fails, the error will be non-nil.

func FetchManufacturers() ([]Manufacturer, error) {
	response, err := http.Get("http://localhost:3000/api/manufacturers")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var manufacturers []Manufacturer
	if err := json.NewDecoder(response.Body).Decode(&manufacturers); err != nil {
		return nil, err
	}
	return manufacturers, nil
}
