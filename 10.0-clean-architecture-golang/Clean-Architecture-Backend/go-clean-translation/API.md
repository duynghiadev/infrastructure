# Translation API Documentation

This document outlines the available API endpoints for the Translation service.

## Base URL

```
http://localhost:8080/v1
```

## Endpoints

### 1. Translate Text

Translates text from one language to another.

- **URL**: `/translate`
- **Method**: `POST`
- **Request Body**:

```json
{
  "original_text": "Hello",
  "source": "en",
  "destination": "vi"
}
```

- **Response**:

```json
{
  "original_text": "Hello",
  "source": "en",
  "destination": "vi",
  "result_text": "Xin chào"
}
```

- **Error Response**:

```json
{
  "error": "error message"
}
```

### 2. Get Translation History

Retrieves the history of all translations.

- **URL**: `/histories`
- **Method**: `GET`
- **Response**:

```json
[
  {
    "original_text": "Hello",
    "source": "en",
    "destination": "vi",
    "result_text": "Xin chào"
  },
  {
    "original_text": "Goodbye",
    "source": "en",
    "destination": "vi",
    "result_text": "Tạm biệt"
  }
]
```

- **Error Response**:

```json
{
  "error": "error message"
}
```

## Language Codes

The service uses standard language codes for the `source` and `destination` parameters. Some common examples:

- `en`: English
- `vi`: Vietnamese
- `es`: Spanish
- `fr`: French
- `de`: German
- `ja`: Japanese
- `ko`: Korean
- `zh`: Chinese

## Error Handling

### HTTP Status Codes

#### 400 Bad Request

Occurs when:

- Missing required fields (`original_text`, `source`, or `destination`)
- Invalid language codes
- Empty text for translation
- Invalid JSON format in request body

Example response:

```json
{
  "error": "Invalid request: missing required field 'original_text'"
}
```

#### 500 Internal Server Error

Occurs when:

- Database connection fails
- Google Translate service is unavailable
- Server encounters an unexpected error
- Memory or processing limitations are reached

Example response:

```json
{
  "error": "Internal server error: translation service unavailable"
}
```

### Error Prevention

1. Always validate input parameters before making API calls
2. Ensure language codes are supported
3. Check that text length is within reasonable limits
4. Handle network timeouts appropriately
5. Implement retry logic for temporary failures

## Notes

1. The translation service uses Google Translate under the hood
2. Translations are cached in the database for faster retrieval of previously translated texts
3. All API responses use UTF-8 encoding to properly handle international characters
