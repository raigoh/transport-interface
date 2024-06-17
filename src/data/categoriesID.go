package data

import (
	"encoding/json"
	"net/http"
	"strconv"
)

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
