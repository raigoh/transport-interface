package data

import (
	"encoding/json"
	"net/http"
	"strconv"
)

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
