# API Documentation

## Base URL

```
http://localhost:8081/api/v1
```

## Authentication Endpoints

### Register User

```
POST /auth/register
```

**Request Body:**

```json
{
  "username": "duy nghia",
  "password": "12345",
  "limit": 200
}
```

**Response:**

- **Status:** 201 Created
- **Body:** JWT Token

### Login

```
POST /auth/login
```

**Request Body:**

```json
{
  "username": "duy nghia",
  "password": "12345"
}
```

**Response:**

- **Status:** 200 OK
- **Body:** JWT Token

## Todo Endpoints

### Create Todo

```
POST /todos
```

**Headers:**

```
Authorization: Bearer {jwt_token}
```

**Request Body:**

```json
{
  "content": "test post 3"
}
```

**Response:**

- **Status:** 201 Created
- **Body:** Created todo object

### Get User's Todos

```
GET /todos/{userId}
```

Example: `/todos/a17ab7c8-d2a6-4af2-a8ae-bc45cdcc65d8`

**Headers:**

```
Authorization: Bearer {jwt_token}
```

**Response:**

- **Status:** 200 OK
- **Body:** Array of todos

### Get All Todos

```
GET /todos
```

**Response:**

- **Status:** 200 OK
- **Body:** Array of all todos

## Error Responses

### Common Error Status Codes

- **400:** Bad Request - Invalid input data
- **401:** Unauthorized - Invalid or missing token
- **403:** Forbidden - Token valid but insufficient permissions
- **404:** Not Found - Resource not found
- **500:** Internal Server Error

### Error Response Format

```json
{
  "error": "Error message description"
}
```

## Notes

- All requests must include the header `Content-Type: application/json`
- Protected endpoints require a JWT token in the `Authorization` header
- Token format: `Bearer {jwt_token}`
- Request body size limit: 2MB
