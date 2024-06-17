package data

import (
	"encoding/json"
	"net/http"
)

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
