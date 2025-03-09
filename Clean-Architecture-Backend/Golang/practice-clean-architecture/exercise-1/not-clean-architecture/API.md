# TODO API Documentation

This document outlines the available endpoints and their usage for the TODO API.

## Base URL

```
http://localhost:8080/v1
```

## Endpoints

### 1. Create a New Todo Item

**Endpoint:** `POST /items`

**Request Body:**

```json
{
  "title": "Your todo item title"
}
```

Note: Status will be automatically set to "Doing"

**Response:**

```json
{
  "data": 1 // Returns the ID of created item
}
```

**Error Response:**

```json
{
  "error": "error message"
}
```

### 2. Get List of Todo Items

**Endpoint:** `GET /items`

**Query Parameters:**

- `page` (optional, default: 1): Page number
- `limit` (optional, default: 10): Number of items per page

**Response:**

```json
{
  "data": [
    {
      "id": 1,
      "title": "Todo item title",
      "status": "Doing",
      "created_at": "2024-03-20T10:00:00Z",
      "updated_at": "2024-03-20T10:00:00Z"
    }
    // ... more items
  ]
}
```

### 3. Get a Todo Item by ID

**Endpoint:** `GET /items/:id`

**Response:**

```json
{
  "data": {
    "id": 1,
    "title": "Todo item title",
    "status": "Doing",
    "created_at": "2024-03-20T10:00:00Z",
    "updated_at": "2024-03-20T10:00:00Z"
  }
}
```

### 4. Update a Todo Item

**Endpoint:** `PUT /items/:id`

**Request Body:**

```json
{
  "title": "Updated title",
  "status": "Done"
}
```

**Response:**

```json
{
  "data": true
}
```

### 5. Delete a Todo Item

**Endpoint:** `DELETE /items/:id`

**Response:**

```json
{
  "data": true
}
```

## Error Handling

All endpoints may return error responses in the following format:

```json
{
  "error": "Description of the error"
}
```

Common HTTP Status Codes:

- 200: Success
- 400: Bad Request
- 404: Not Found
- 500: Internal Server Error
