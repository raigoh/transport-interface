package data

import (
	"encoding/json"
	"net/http"
)

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
