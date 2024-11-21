package data

import (
	"encoding/json"
	"net/http"
)

// FetchCategories retrieves the list of categories from the API.
// It sends a GET request to the /api/categories endpoint and decodes the JSON response
// into a slice of Category structs.
//
// Returns:
//   - ([]Category, error): A slice of Category structs and an error.
//     If the request or decoding fails, the error will be non-nil.

func FetchCategories() ([]Category, error) {
	response, err := http.Get("http://localhost:3000/api/categories")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var categories []Category
	if err := json.NewDecoder(response.Body).Decode(&categories); err != nil {
		return nil, err
	}
	return categories, nil
}
