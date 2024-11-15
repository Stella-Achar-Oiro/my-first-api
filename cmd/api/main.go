package main

import (
    "log"
    "my-first-api/internal/todo"
    "my-first-api/internal/transport"
)

func main() {
    todoService := todo.NewService()
    server := transport.NewServer(todoService)

    log.Fatal(server.Serve())
}