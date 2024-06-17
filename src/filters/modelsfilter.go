package filters

import (
	"cars/src/data"
	"strconv"
	"strings"
)

func FetchModelsFiltered(manufacturer, minYear, maxYear, category, searchQuery string) ([]data.CarModel, error) {
	models, err := data.FetchModels()
	if err != nil {
		return nil, err
	}

	var filteredModels []data.CarModel

	// Convert category and years to integers if provided
	var categoryID, minYearInt, maxYearInt int
	if category != "" {
		categoryID, err = strconv.Atoi(category)
		if err != nil {
			return nil, err
		}
	}
	if minYear != "" {
		minYearInt, err = strconv.Atoi(minYear)
		if err != nil {
			return nil, err
		}
	}
	if maxYear != "" {
		maxYearInt, err = strconv.Atoi(maxYear)
		if err != nil {
			return nil, err
		}
	}

	for _, model := range models {
		// Fetch manufacturer details for each model
		manufacturerDetails, err := data.FetchManufacturerByID(model.ManufacturerID)
		if err != nil {
			return nil, err
		}

		// Check if the model matches the filters
		matchesManufacturer := (manufacturer == "" || manufacturerDetails.Name == manufacturer)
		matchesCategory := (category == "" || model.CategoryID == categoryID)
		matchesMinYear := (minYear == "" || model.Year >= minYearInt)
		matchesMaxYear := (maxYear == "" || model.Year <= maxYearInt)
		matchesSearchQuery := (searchQuery == "" || strings.Contains(strings.ToLower(model.Name), strings.ToLower(searchQuery)))

		if matchesManufacturer && matchesCategory && matchesMinYear && matchesMaxYear && matchesSearchQuery {
			filteredModels = append(filteredModels, model)
		}
	}

	return filteredModels, nil
}
