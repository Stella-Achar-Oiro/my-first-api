package todo

import (
	"errors"
	"fmt"
	"my-first-api/internal/models"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Service handles todo business logic
type Service struct {
	todos []models.Todo
	mu    sync.RWMutex
}

// NewService creates a new todo service
func NewService() *Service {
	return &Service{
		todos: make([]models.Todo, 0),
	}
}

// Add creates a new todo
func (s *Service) Add(input models.TodoInput) error {
	if input.Task == "" {
		return errors.New("task cannot be empty")
	}

	if input.Priority == "" {
		input.Priority = models.PriorityMedium
	}

	if !models.ValidPriority(input.Priority) {
		return fmt.Errorf("invalid priority: %s", input.Priority)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// Check for duplicates
	for _, t := range s.todos {
		if strings.EqualFold(t.Task, input.Task) {
			return errors.New("todo item already exists")
		}
	}

	now := time.Now()
	todo := models.Todo{
		ID:        uuid.New().String(),
		Task:      input.Task,
		Status:    models.StatusToBeStarted,
		Priority:  input.Priority,
		DueDate:   input.DueDate,
		CreatedAt: now,
		UpdatedAt: now,
	}

	s.todos = append(s.todos, todo)
	return nil
}

// UpdateStatus updates a todo's status
func (s *Service) UpdateStatus(item, status string) error {
	if !models.ValidStatus(status) {
		return fmt.Errorf("invalid status: %s", status)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	for i := range s.todos {
		if strings.EqualFold(s.todos[i].Task, item) {
			s.todos[i].Status = status
			s.todos[i].UpdatedAt = time.Now()

			if status == models.StatusCompleted {
				s.todos[i].CompletedAt = time.Now()
			}
			return nil
		}
	}

	return errors.New("todo item not found")
}

// Delete removes a todo
func (s *Service) Delete(task string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, t := range s.todos {
		if strings.EqualFold(t.Task, task) {
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo item not found")
}

// GetByPriority returns todos filtered by priority
func (s *Service) GetByPriority(priority string) ([]models.Todo, error) {
	if !models.ValidPriority(priority) {
		return nil, fmt.Errorf("invalid priority: %s", priority)
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []models.Todo
	for _, todo := range s.todos {
		if strings.EqualFold(todo.Priority, priority) {
			result = append(result, todo)
		}
	}
	return result, nil
}

// GetByStatus returns todos filtered by status
func (s *Service) GetByStatus(status string) ([]models.Todo, error) {
	if !models.ValidStatus(status) {
		return nil, fmt.Errorf("invalid status: %s", status)
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []models.Todo
	for _, todo := range s.todos {
		if todo.Status == status {
			result = append(result, todo)
		}
	}
	return result, nil
}

// Search searches todos based on a query string
func (s *Service) Search(query string) []models.Todo {
	if query == "" {
		return nil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	var results []models.Todo
	for _, todo := range s.todos {
		if strings.Contains(
			strings.ToLower(todo.Task),
			strings.ToLower(query),
		) {
			results = append(results, todo)
		}
	}
	return results
}

// GetAll returns all todos
func (s *Service) GetAll() []models.Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// Return a copy to prevent external modifications
	todos := make([]models.Todo, len(s.todos))
	copy(todos, s.todos)
	return todos
}
