# Todo API Documentation

## Overview
This API provides endpoints for managing todo items, including creation, retrieval, updates, and deletion of todos. The API follows RESTful principles and uses JSON for request and response bodies.

## Base URL
```
http://localhost:8080
```

## Authentication
Currently, the API is open and doesn't require authentication. Future versions will implement JWT-based authentication.

## Common Response Codes
- `200 OK`: Request successful
- `201 Created`: Resource created successfully
- `400 Bad Request`: Invalid request payload
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource already exists
- `500 Internal Server Error`: Server error

## Endpoints

### Health Check

#### GET /health
Check the API's health status.

**Response**
```json
{
    "status": "ok",
    "time": "2024-11-15T10:00:00Z"
}
```

### Todo Management

#### GET /todo
Retrieve all todos.

**Response**
```json
[
    {
        "id": "1",
        "task": "Complete project",
        "status": "IN_PROGRESS",
        "priority": "high",
        "dueDate": "2024-12-31T23:59:59Z",
        "createdAt": "2024-11-15T10:00:00Z",
        "updatedAt": "2024-11-15T10:00:00Z"
    }
]
```

#### POST /todo
Create a new todo.

**Request Body**
```json
{
    "task": "Complete project",
    "priority": "high",
    "dueDate": "2024-12-31T23:59:59Z"
}
```

**Response (201 Created)**
```json
{
    "message": "Todo created successfully",
    "task": "Complete project"
}
```

#### PUT /todo/status
Update a todo's status.

**Request Body**
```json
{
    "item": "Complete project",
    "status": "IN_PROGRESS"
}
```

**Valid Status Values**
- TO_BE_STARTED
- IN_PROGRESS
- COMPLETED

**Response**
```json
{
    "message": "Todo status updated successfully",
    "item": "Complete project",
    "status": "IN_PROGRESS"
}
```

#### DELETE /todo
Delete a todo.

**Request Body**
```json
{
    "item": "Complete project"
}
```

**Response**
```json
{
    "message": "Todo deleted successfully",
    "item": "Complete project"
}
```

### Priority Management

#### GET /todo/priority/{priority}
Get todos filtered by priority level.

**Parameters**
- `priority` (path): Priority level (low, medium, high)

**Response**
```json
[
    {
        "id": "1",
        "task": "Complete project",
        "status": "IN_PROGRESS",
        "priority": "high",
        "dueDate": "2024-12-31T23:59:59Z",
        "createdAt": "2024-11-15T10:00:00Z",
        "updatedAt": "2024-11-15T10:00:00Z"
    }
]
```

### Search

#### GET /search
Search todos based on a query string.

**Parameters**
- `q` (query): Search term

**Example**
```
GET /search?q=project
```

**Response**
```json
[
    {
        "id": "1",
        "task": "Complete project",
        "status": "IN_PROGRESS",
        "priority": "high",
        "dueDate": "2024-12-31T23:59:59Z",
        "createdAt": "2024-11-15T10:00:00Z",
        "updatedAt": "2024-11-15T10:00:00Z"
    }
]
```

## Data Models

### Todo
```json
{
    "id": "string",
    "task": "string",
    "status": "string",
    "priority": "string",
    "dueDate": "string (ISO 8601)",
    "createdAt": "string (ISO 8601)",
    "updatedAt": "string (ISO 8601)",
    "completedAt": "string (ISO 8601)"
}
```

### Todo Input
```json
{
    "task": "string (required, 3-100 chars)",
    "priority": "string (low, medium, high)",
    "dueDate": "string (ISO 8601)"
}
```

## Error Handling

### Error Response Format
```json
{
    "error": "Error message description"
}
```

### Common Error Messages
- "Task cannot be empty"
- "Todo item already exists"
- "Invalid status"
- "Todo item not found"
- "Invalid request payload"

## Rate Limiting
The API currently doesn't implement rate limiting. Future versions will include rate limiting headers:
- X-RateLimit-Limit
- X-RateLimit-Remaining
- X-RateLimit-Reset

## Pagination
For endpoints that return lists (GET /todo, GET /search), pagination will be implemented in future versions using:
- page (query parameter)
- limit (query parameter)
- total (response header)

## Best Practices
1. Always include Content-Type header in requests
2. Handle HTTP 429 (Too Many Requests) in clients
3. Implement exponential backoff for failed requests
4. Cache GET responses when appropriate
5. Use appropriate HTTP methods for operations

## Future Enhancements
1. Authentication and Authorization
2. Pagination
3. Rate Limiting
4. Batch Operations
5. Webhooks for Status Changes
6. Rich Text Descriptions
7. File Attachments
8. Categories and Tags