# Clean Architecture Guide

Clean Architecture is a software design philosophy that separates the concerns of a software system into layers, making it more maintainable, testable, and independent of external frameworks or tools. This guide will demonstrate Clean Architecture implementation in different programming languages, starting with Go.

## Table of Contents

- [Overview](#overview)
- [Core Principles](#core-principles)
- [Implementation in Go](#implementation-in-go)
- [Implementation in Other Languages](#implementation-in-other-languages)
- [Best Practices](#best-practices)

## Overview

Clean Architecture, proposed by Robert C. Martin (Uncle Bob), consists of four main layers:

1. **Entities** (Enterprise Business Rules)
2. **Use Cases** (Application Business Rules)
3. **Interface Adapters** (Controllers, Presenters, Gateways)
4. **Frameworks & Drivers** (Web, UI, Database, External Interfaces)

![Clean Architecture Diagram](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

## Core Principles

1. **Independence of Frameworks**: The architecture doesn't depend on frameworks
2. **Testability**: Business rules can be tested without external elements
3. **Independence of UI**: The UI can change without changing the system
4. **Independence of Database**: Business rules aren't bound to the database
5. **Independence of External Agency**: Business rules don't know anything about external interfaces

## Implementation in Go

### Project Structure

```
├── cmd
│   └── main.go
├── internal
│   ├── domain
│   │   └── entity
│   ├── usecase
│   │   └── repository
│   └── interface
│       └── presenter
└── pkg
    └── common
```

### Example Implementation

#### 1. Entity (Domain Layer)

```go
// internal/domain/entity/user.go
package entity

type User struct {
    ID       string
    Name     string
    Email    string
    Password string
}

func NewUser(name, email, password string) *User {
    return &User{
        Name:     name,
        Email:    email,
        Password: password,
    }
}
```

#### 2. Repository Interface (Use Case Layer)

```go
// internal/usecase/repository/user_repository.go
package repository

type UserRepository interface {
    Create(user *entity.User) error
    GetByID(id string) (*entity.User, error)
    Update(user *entity.User) error
    Delete(id string) error
}
```

#### 3. Use Case (Application Layer)

```go
// internal/usecase/user_usecase.go
package usecase

type UserUseCase struct {
    userRepo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
    return &UserUseCase{
        userRepo: repo,
    }
}

func (uc *UserUseCase) CreateUser(name, email, password string) error {
    user := entity.NewUser(name, email, password)
    return uc.userRepo.Create(user)
}
```

#### 4. Interface Adapter (Controller)

```go
// internal/interface/http/handler/user_handler.go
package handler

type UserHandler struct {
    userUseCase *usecase.UserUseCase
}

func NewUserHandler(useCase *usecase.UserUseCase) *UserHandler {
    return &UserHandler{
        userUseCase: useCase,
    }
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    // Handle HTTP request
    var input struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.userUseCase.CreateUser(input.Name, input.Email, input.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
```

## Implementation in Other Languages

### TypeScript/Node.js Example

#### Entity

```typescript
// src/domain/entities/User.ts
export class User {
  constructor(
    public readonly id: string,
    public readonly name: string,
    public readonly email: string,
    private password: string
  ) {}

  validatePassword(password: string): boolean {
    return this.password === password;
  }
}
```

#### Use Case

```typescript
// src/usecases/CreateUser.ts
export class CreateUserUseCase {
  constructor(private userRepository: UserRepository) {}

  async execute(userData: CreateUserDTO): Promise<User> {
    const user = new User(
      generateId(),
      userData.name,
      userData.email,
      userData.password
    );

    await this.userRepository.save(user);
    return user;
  }
}
```

### Python Example

#### Entity

```python
# domain/entities/user.py
from dataclasses import dataclass
from datetime import datetime

@dataclass
class User:
    id: str
    name: str
    email: str
    created_at: datetime

    @staticmethod
    def create(name: str, email: str) -> 'User':
        return User(
            id=str(uuid.uuid4()),
            name=name,
            email=email,
            created_at=datetime.now()
        )
```

#### Use Case

```python
# usecases/create_user.py
class CreateUserUseCase:
    def __init__(self, user_repository: UserRepository):
        self.user_repository = user_repository

    def execute(self, name: str, email: str) -> User:
        user = User.create(name, email)
        self.user_repository.save(user)
        return user
```

## Best Practices

1. **Dependency Rule**: Dependencies should point inward. Inner layers should not know about outer layers.
2. **Interface Segregation**: Keep interfaces small and focused.
3. **Single Responsibility**: Each component should have one reason to change.
4. **Dependency Injection**: Use dependency injection to maintain loose coupling.
5. **Testing Strategy**:

   - Unit tests for Use Cases
   - Integration tests for Repositories
   - End-to-end tests for Controllers

## Common Mistakes to Avoid

1. Mixing business logic with framework code
2. Allowing domain entities to depend on external libraries
3. Tight coupling between layers
4. Breaking the dependency rule
5. Not using interfaces for dependency inversion

## Conclusion

Clean Architecture provides a robust foundation for building maintainable and scalable applications. By following these principles and patterns, you can create systems that are:

- Easy to maintain and modify
- Highly testable
- Framework independent
- Database independent
- UI independent

Remember that Clean Architecture is a guideline, not a strict ruleset. Adapt it to your specific needs while maintaining the core principles.
