package transport

import (
    "encoding/json"
    "my-first-api/internal/models"
    "my-first-api/internal/todo"
    "net/http"
    "time"
)

type Handlers struct {
    todoService *todo.Service
}

func NewHandlers(ts *todo.Service) *Handlers {
    return &Handlers{
        todoService: ts,
    }
}

func (h *Handlers) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
    respondJSON(w, http.StatusOK, map[string]string{
        "status": "ok",
        "time":   time.Now().Format(time.RFC3339),
    })
}

func (h *Handlers) HandleGetAllTodos(w http.ResponseWriter, r *http.Request) {
    todos := h.todoService.GetAll()
    respondJSON(w, http.StatusOK, todos)
}

func (h *Handlers) HandleCreateTodo(w http.ResponseWriter, r *http.Request) {
    var input models.TodoInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        respondError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if err := h.todoService.Add(input); err != nil {
        if err.Error() == "todo item already exists" {
            respondError(w, http.StatusConflict, err.Error())
            return
        }
        respondError(w, http.StatusBadRequest, err.Error())
        return
    }

    respondJSON(w, http.StatusCreated, map[string]string{
        "message": "Todo created successfully",
        "task":    input.Task,
    })
}

func (h *Handlers) HandleUpdateTodoStatus(w http.ResponseWriter, r *http.Request) {
    var input models.StatusUpdate
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        respondError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if err := h.todoService.UpdateStatus(input.Item, input.Status); err != nil {
        if err.Error() == "todo item not found" {
            respondError(w, http.StatusNotFound, err.Error())
            return
        }
        respondError(w, http.StatusBadRequest, err.Error())
        return
    }

    respondJSON(w, http.StatusOK, map[string]string{
        "message": "Todo status updated successfully",
        "item":    input.Item,
        "status":  input.Status,
    })
}

func (h *Handlers) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Item string `json:"item"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        respondError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if err := h.todoService.Delete(input.Item); err != nil {
        if err.Error() == "todo item not found" {
            respondError(w, http.StatusNotFound, err.Error())
            return
        }
        respondError(w, http.StatusBadRequest, err.Error())
        return
    }

    respondJSON(w, http.StatusOK, map[string]string{
        "message": "Todo deleted successfully",
        "item":    input.Item,
    })
}

func (h *Handlers) HandleSearchTodos(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("q")
    if query == "" {
        respondError(w, http.StatusBadRequest, "Search query is required")
        return
    }

    results := h.todoService.Search(query)
    respondJSON(w, http.StatusOK, results)
}

func (h *Handlers) HandleGetTodosByPriority(w http.ResponseWriter, r *http.Request) {
    priority := r.PathValue("priority")
    if priority == "" {
        respondError(w, http.StatusBadRequest, "Priority is required")
        return
    }

    todos, err := h.todoService.GetByPriority(priority)
    if err != nil {
        respondError(w, http.StatusBadRequest, err.Error())
        return
    }

    respondJSON(w, http.StatusOK, todos)
}

func (h *Handlers) HandleGetTodosByStatus(w http.ResponseWriter, r *http.Request) {
    status := r.PathValue("status")
    if status == "" {
        respondError(w, http.StatusBadRequest, "Status is required")
        return
    }

    todos, err := h.todoService.GetByStatus(status)
    if err != nil {
        respondError(w, http.StatusBadRequest, err.Error())
        return
    }

    respondJSON(w, http.StatusOK, todos)
}

// Helper functions for consistent response handling
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
    response, err := json.Marshal(payload)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    w.Write(response)
}

func respondError(w http.ResponseWriter, code int, message string) {
    respondJSON(w, code, map[string]string{"error": message})
}