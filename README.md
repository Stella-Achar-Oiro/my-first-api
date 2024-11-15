# Go Todo API

A professional-grade RESTful API for task management built with Go. This project demonstrates best practices in API design, error handling, and code organization.

## 🌟 Features

- RESTful API endpoints for todo management
- Priority levels (high, medium, low)
- Status tracking (TO_BE_STARTED, IN_PROGRESS, COMPLETED)
- Search functionality
- Due dates and timestamps
- Input validation
- Error handling
- CORS support
- Logging middleware
- API documentation with Swagger
- Comprehensive test suite

## 🛠️ Technology Stack

- Go 1.22+
- Standard library for HTTP server
- UUID for unique identifiers
- VSCode REST Client for testing

## 🚀 Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/my-first-api.git
cd my-first-api
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`

## 📖 API Documentation

Detailed API documentation is available in:
- [API Documentation](docs/api.md)
- [Swagger Specification](docs/swagger.yaml)

### Quick Endpoint Reference:

```bash
# Health Check
GET /health

# Todo Operations
GET    /todo                    # Get all todos
POST   /todo                    # Create a todo
PUT    /todo/status            # Update todo status
DELETE /todo                    # Delete a todo

# Filtering & Search
GET /todo/priority/{priority}   # Get todos by priority
GET /todo/status/{status}      # Get todos by status
GET /search?q={query}          # Search todos
```

## 🧪 Testing

Use the provided `api.http` file with VSCode REST Client to test all endpoints:

1. Install REST Client extension in VSCode
2. Open `api.http`
3. Click "Send Request" above any request to test

## 📁 Project Structure

```
my-first-api/
├── api.http                 # API test file
├── cmd/
│   └── api/
│       └── main.go         # Application entry point
├── config/
│   └── config.go           # Configuration management
├── internal/
│   ├── middleware/         # HTTP middleware
│   ├── models/             # Data models
│   ├── todo/               # Business logic
│   └── transport/          # HTTP handlers and routing
├── docs/                   # Documentation
└── go.mod                  # Go modules file
```

## 🔄 Future Enhancements

- [ ] Database integration (PostgreSQL)
- [ ] User authentication
- [ ] Docker support
- [ ] CI/CD pipeline
- [ ] Rate limiting
- [ ] Caching layer
- [ ] Metrics and monitoring

## 💡 Contributing

Feel free to submit issues and enhancement requests!

## 📝 License

This project is open-source and available under the MIT License.