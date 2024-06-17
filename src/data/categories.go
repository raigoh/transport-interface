package data

import (
	"encoding/json"
	"net/http"
)

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
