# Golang Transaction Project

## Project Structure

This project demonstrates a multi-layer architecture using Golang and GORM for database transactions.

### Main Components

- **Model**: Handles database connections and entity definitions
- **Route**: Contains API route configurations
- **Controller**: (if exists) Handles business logic
- **Repository**: (if exists) Manages data access operations

## Setup and Configuration

### Database Connection

The project uses GORM for database operations. Database connection is established in the `model` package.

### Main Application

The entry point of the application (`main.go`):

```go
package main

import (
    "golang-transaction/model"
    "golang-transaction/route"
)

func main() {
    db, _ := model.DBConnection()
    route.SetupRoutes(db)
}
```

## Running the Application

1. Ensure you have Go installed on your system
2. Set up your database configuration
3. Run the application:

```bash
go run main.go
```
