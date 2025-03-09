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

## Error Codes

- `400`: Bad Request - Invalid input parameters
- `500`: Internal Server Error - Server-side error

## Notes

1. The translation service uses Google Translate under the hood
2. Translations are cached in the database for faster retrieval of previously translated texts
3. All API responses use UTF-8 encoding to properly handle international characters
