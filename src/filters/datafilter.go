package filters

import "cars/src/data"

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
