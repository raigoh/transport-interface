package filters

import "cars/src/data"

// FetchFilterData retrieves filter data for car models from the API.
// It fetches manufacturers, car models, and categories, then extracts unique years from the car models.
// The function returns a FilterData struct containing lists of manufacturer names, years, and category names.
//
// Returns:
//   - (data.FilterData, error): A FilterData struct and an error.
//     If any of the API requests fail, the error will be non-nil.

func FetchFilterData() (data.FilterData, error) {
	manufacturers, err := data.FetchManufacturers()
	if err != nil {
		return data.FilterData{}, err
	}

	models, err := data.FetchModels()
	if err != nil {
		return data.FilterData{}, err
	}

	categories, err := data.FetchCategories()
	if err != nil {
		return data.FilterData{}, err
	}

	years := make(map[int]bool)
	for _, model := range models {
		years[model.Year] = true
	}

	yearList := make([]int, 0, len(years))
	for year := range years {
		yearList = append(yearList, year)
	}

	manufacturerNames := make([]string, len(manufacturers))
	for i, manufacturer := range manufacturers {
		manufacturerNames[i] = manufacturer.Name
	}

	categoryNames := make([]string, len(categories))
	for i, category := range categories {
		categoryNames[i] = category.Name
	}

	return data.FilterData{
		Manufacturers: manufacturerNames,
		Years:         yearList,
		Categories:    categoryNames,
	}, nil
}
