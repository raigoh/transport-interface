package data

// Define a struct to represent a car manufacturer
type Manufacturer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Country      string `json:"country"`
	FoundingYear int    `json:"foundingYear"`
}

// Define a struct to represent a car category
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Define a struct to represent a car model
type CarModel struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	ManufacturerID int    `json:"manufacturerId"`
	CategoryID     int    `json:"categoryId"`
	Year           int    `json:"year"`
	Specifications struct {
		Engine       string `json:"engine"`
		Horsepower   int    `json:"horsepower"`
		Transmission string `json:"transmission"`
		Drivetrain   string `json:"drivetrain"`
	} `json:"specifications"`
	Image               string       `json:"image"`
	ManufacturerDetails Manufacturer `json:"manufacturerDetails"`
	CategoryDetails     Category     `json:"categoryDetails"`
}

type FilterData struct {
	Manufacturers []string
	Years         []int
	Categories    []string
}
