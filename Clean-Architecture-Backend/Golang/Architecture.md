# Golang Clean Architecture Guide

## Overview

Clean Architecture in Golang emphasizes separation of concerns and dependency rules, making the codebase more maintainable, testable, and scalable.

## Project Structure

## Layers Description

### 1. Domain Layer

- Contains business logic and entities
- No dependencies on external layers
- Pure Go structs and interfaces

```go
// Example domain entity
type User struct {
    ID        string
    Username  string
    Email     string
    CreatedAt time.Time
}

// Domain repository interface
type UserRepository interface {
    GetByID(id string) (*User, error)
    Save(user *User) error
}
```

### 2. Use Case Layer

- Contains application business rules
- Implements use case interfaces
- Depends on domain layer

```go
type UserUseCase interface {
    CreateUser(username, email string) (*User, error)
    GetUser(id string) (*User, error)
}
```

### 3. Repository Layer

- Implements data storage interfaces
- Handles database operations
- Converts between domain and data models

### 4. Delivery Layer

- Handles HTTP/gRPC/CLI interfaces
- Contains controllers/handlers
- Manages request/response cycles

## SOLID Principles in Golang

### 1. Single Responsibility Principle (SRP)

- Each struct/interface has one responsibility
- Example:

```go
// Good: Single responsibility
type UserService struct {
    repo UserRepository
}

func (s *UserService) CreateUser(user *User) error {
    return s.repo.Save(user)
}
```

### 2. Open/Closed Principle (OCP)

- Open for extension, closed for modification
- Use interfaces for flexibility

```go
// Interface allowing different implementations
type Storage interface {
    Save(data interface{}) error
    Get(id string) (interface{}, error)
}

// Multiple implementations without modifying interface
type MySQLStorage struct{}
type MongoStorage struct{}
```

### 3. Liskov Substitution Principle (LSP)

- Subtypes must be substitutable for their base types
- Example:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Both implementations are substitutable
type FileReader struct{}
type NetworkReader struct{}
```

### 4. Interface Segregation Principle (ISP)

- Keep interfaces small and focused
- Example:

```go
// Good: Segregated interfaces
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Instead of one large interface
type ReadWriter interface {
    Reader
    Writer
}
```

### 5. Dependency Inversion Principle (DIP)

- High-level modules shouldn't depend on low-level modules
- Both should depend on abstractions
- Example:

```go
// High-level module depends on abstraction
type UserService struct {
    repo UserRepository // interface
}

// Low-level module implements abstraction
type MySQLUserRepository struct{}
```

## Best Practices

1. **Dependency Injection**

   - Use constructor injection
   - Implement interfaces for testing
   - Use DI containers when necessary

2. **Error Handling**

   - Create custom error types
   - Use error wrapping
   - Implement error handling middleware

3. **Configuration**

   - Use environment variables
   - Implement config structs
   - Separate config by environment

4. **Testing**

   - Unit tests for business logic
   - Integration tests for repositories
   - Mock interfaces for dependencies

## Example Implementation

```go
// Domain
type User struct {
    ID   string
    Name string
}

// Repository Interface
type UserRepository interface {
    Save(user *User) error
    GetByID(id string) (*User, error)
}

// Use Case
type UserUseCase struct {
    repo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
    return &UserUseCase{repo: repo}
}

// Delivery
type UserHandler struct {
    useCase *UserUseCase
}

func NewUserHandler(useCase *UserUseCase) *UserHandler {
    return &UserHandler{useCase: useCase}
}
```

## Conclusion

Following Clean Architecture and SOLID principles in Golang leads to:

- Maintainable and testable code
- Clear separation of concerns
- Easy to extend functionality
- Better dependency management
- Improved code organization

Remember to adapt these principles based on your project's specific needs while maintaining the core concepts of Clean Architecture.
