# Data Access Object (DAO) in Golang

## Introduction

The **Data Access Object (DAO)** pattern is a widely used approach in software development to abstract and encapsulate database access logic. In Golang, where simplicity and maintainability are key principles, the DAO pattern helps structure applications by separating database operations from business logic.

## Why Use DAO in Golang?

1. **Separation of Concerns** – DAO isolates data access logic from the rest of the application.
2. **Flexibility** – Easily swap databases (e.g., MySQL to PostgreSQL) by implementing different DAO structures.
3. **Testability** – Using interfaces makes it easier to mock the database layer for testing.
4. **Readability and Maintainability** – Encourages clear and structured code organization.

## DAO Structure in Golang

A typical DAO implementation in Go includes:

- **Models** : Structs representing database entities.
- **Interfaces** : Define contract methods for data access.
- **Repositories** : Implement interfaces to interact with the database.
- **Services** : Business logic layer interacting with DAO.

### Example DAO Implementation

#### 1. Define the Model (models/user.go)

```go
package models

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

#### 2. Define the DAO Interface (repositories/user_repository.go)

```go
package repositories

import "github.com/your_project/models"

type UserRepository interface {
    GetUserByID(id int) (*models.User, error)
    CreateUser(user *models.User) error
}
```

#### 3. Implement DAO for MySQL (repositories/mysql_user_repository.go)

```go
package repositories

import (
    "database/sql"
    "github.com/your_project/models"
)

type MySQLUserRepository struct {
    DB *sql.DB
}

func (repo *MySQLUserRepository) GetUserByID(id int) (*models.User, error) {
    user := &models.User{}
    err := repo.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).
        Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    return user, nil
}

func (repo *MySQLUserRepository) CreateUser(user *models.User) error {
    _, err := repo.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
    return err
}
```

#### 4. Use DAO in a Service (services/user_service.go)

```go
package services

import (
    "github.com/your_project/models"
    "github.com/your_project/repositories"
)

type UserService struct {
    UserRepo repositories.UserRepository
}

func (s *UserService) GetUser(id int) (*models.User, error) {
    return s.UserRepo.GetUserByID(id)
}
```

#### 5. Dependency Injection in main.go

```go
package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
    "github.com/your_project/repositories"
    "github.com/your_project/services"
)

func main() {
    db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    userRepo := &repositories.MySQLUserRepository{DB: db}
    userService := &services.UserService{UserRepo: userRepo}

    user, err := userService.GetUser(1)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("User:", user)
}
```

## Conclusion

By using the DAO pattern in Go, we achieve:

- **Loose coupling** between business logic and database operations.
- **Scalability** by easily switching databases.
- **Better testability** by mocking the DAO interface in unit tests.

This pattern is widely used in Go backend development for clean, maintainable, and testable applications.
