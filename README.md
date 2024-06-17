# Cars Inventory Management System

Welcome to the Cars Inventory Management System! This application allows you to manage and view car models, manufacturers, and categories. It also provides functionalities for filtering and comparing different car models.

## Table of Contents

- [Features](#features)
- [Getting Started](#gettingstarted)
- [Installation](#installation)
- [Running the Application](#application)
- [API Endpoints](#api-endpoints)
- [Middleware](#middleware)
- [Templates](#templates)
- [Contributing](#contributing)
- [Contact](#contact)

## Features

- Fetch and display car categories
- Fetch and display car manufacturers
- Fetch and display car models
- View details of a specific car model, including its manufacturer and category details
- Filter car models by manufacturer, year, category, and search query
- Compare multiple car models
- Health check endpoint for monitoring
- Experimental endpoint for testing purposes
- Error handling and recovery middleware
- Logging middleware
- Rate limiting middleware
- Caching middleware for improved performance

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- A running API server that provides the following endpoints:
  - GET /api/categories
  - GET /api/categories/{id}
  - GET /api/manufacturers
  - GET /api/manufacturers/{id}
  - GET /api/models
  - GET /api/models/{id}

## Installation

1. Clone the repository:

```bash
git clone https://gitea.koodsisu.fi/raigohoim/cars.git
cd cars
```

2. Install dependencies:

```bash
go mod tidy
```

## Running the Application

1.Start the API server:

```bash
cd api
make run
```

2. Navigate to the cars directory:

```bash
cd cars
```

3. Run the application:

```bash
go run main.go
```

4. Open your web browser and navigate to http://localhost:8080 or make a ctrl + click on terminal to http://localhost:8080.

## API Endpoints

- GET /api/categories: Fetches all car categories.
- GET /api/categories/{id}: Fetches a specific car category by ID.
- GET /api/manufacturers: Fetches all car manufacturers.
- GET /api/manufacturers/{id}: Fetches a specific car manufacturer by ID.
- GET /api/models: Fetches all car models.
- GET /api/models/{id}: Fetches a specific car model by ID.

## Middleware

- LoggingMiddleware: Logs incoming requests with method, URI, remote address, and processing time.
- RateLimitingMiddleware: Limits the number of requests from a single IP address.
- CachingMiddleware: Caches GET responses to reduce load and improve performance.
- WithRecovery: Recovers from panics and renders a friendly error page.

## Templates

- index.html: Main page that displays the list of car models and filter options.
- car.html: Detailed view of a single car model.
- compare.html: Comparison view for selected car models.
- error.html: Custom error page displayed in case of server errors.

## Contributing

Contributions are welcome! If you find any issues or want to add new features, please open an issue or submit a pull request.

## Contact

For questions or feedback, please contact discord Vikationu#4963.
