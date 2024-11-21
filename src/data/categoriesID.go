package data

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// FetchCategoryByID retrieves a category by its ID from the API.
// It sends a GET request to the /api/categories/{id} endpoint and decodes the JSON response
// into a Category struct.
//
// Parameters:
//   - id (int): The ID of the category to retrieve.
//
// Returns:
//   - (Category, error): A Category struct and an error.
//     If the request or decoding fails, the error will be non-nil.

func FetchCategoryByID(id int) (Category, error) {
	response, err := http.Get("http://localhost:3000/api/categories/" + strconv.Itoa(id))
	if err != nil {
		return Category{}, err
	}
	defer response.Body.Close()

	var category Category
	if err := json.NewDecoder(response.Body).Decode(&category); err != nil {
		return Category{}, err
	}

	return category, nil
}
