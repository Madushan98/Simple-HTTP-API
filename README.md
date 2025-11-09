# Simple HTTP API - GoLang

## Overview

A simple Go web service with a single `/hello-world` endpoint that responds based on the first letter of the provided `name` query parameter.

## Requirements

- Go 1.18 or higher

## Run the application

```bash
go run cmd/app/main.go
```

## Endpoints

### GET /hello-world

Returns a greeting message based on the provided name.

**Parameters:**

- `name` (string, required): Name to generate greeting for

**Responses:**

| Condition | Status Code | Response Body |
|-----------|-------------|---------------|
| Name starts with A-M (case-insensitive) | 200 OK | `{ "message": "Hello {name}" }` |
| Name starts with N-Z (case-insensitive) | 400 Bad Request | `{ "error": "Invalid Input" }` |
| Name missing or empty or invalid(name start with non-alphabetic character) | 400 Bad Request | `{ "error": "Invalid Input" }` |

**Example:**

```bash
GET /hello-world?name=Alice
```

```json
{
    "message": "Hello Alice"
}
```

## Run the tests

```bash
go test ./test/... 
```

## Assumptions

- Names starting with non-alphabetic characters are considered invalid.
- Case-insensitive comparison (A–M and a–m are both valid).
- Accented letters (e.g., é, ü, ñ) are considered invalid — only A–Z and a–z are allowed.
