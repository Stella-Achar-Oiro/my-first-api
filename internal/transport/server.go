package transport

import (
    "fmt"
    "my-first-api/internal/todo"
    "net/http"
    "time"
)

type Server struct {
    router *Router
    server *http.Server
}

func NewServer(todoService *todo.Service) *Server {
    router := NewRouter(todoService)
    
    return &Server{
        router: router,
        server: &http.Server{
            Addr:         ":8080",
            Handler:      router,
            ReadTimeout:  10 * time.Second,
            WriteTimeout: 10 * time.Second,
        },
    }
}

func (s *Server) Serve() error {
    fmt.Println("Server starting on :8080")
    return s.server.ListenAndServe()
}
