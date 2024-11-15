package transport

import (
    "my-first-api/internal/middleware"
    "my-first-api/internal/todo"
    "net/http"
)

type Router struct {
    *http.ServeMux
    handlers *Handlers
}

func NewRouter(todoService *todo.Service) *Router {
    router := &Router{
        ServeMux: http.NewServeMux(),
        handlers: NewHandlers(todoService),
    }
    router.setupRoutes()
    return router
}

func (r *Router) setupRoutes() {
    // Wrap the entire router with common middleware
    r.Handle("/", middleware.CORSMiddleware([]string{"*"})(
        middleware.LoggingMiddleware(r.routes())))
}

func (r *Router) routes() http.Handler {
    mux := http.NewServeMux()

    // Health check
    mux.HandleFunc("GET /health", r.handlers.HandleHealthCheck)

    // Todo routes
    mux.HandleFunc("GET /todo", r.handlers.HandleGetAllTodos)
    mux.HandleFunc("POST /todo", r.handlers.HandleCreateTodo)
    mux.HandleFunc("PUT /todo/status", r.handlers.HandleUpdateTodoStatus)
    mux.HandleFunc("DELETE /todo", r.handlers.HandleDeleteTodo)
    
    // Priority and Status routes
    mux.HandleFunc("GET /todo/priority/{priority}", r.handlers.HandleGetTodosByPriority)
    mux.HandleFunc("GET /todo/status/{status}", r.handlers.HandleGetTodosByStatus)
    
    // Search route
    mux.HandleFunc("GET /search", r.handlers.HandleSearchTodos)

    return mux
}