package models

import (
    "time"
)

// Todo represents a todo item in the system
type Todo struct {
    ID          string    `json:"id"`
    Task        string    `json:"task"`
    Status      string    `json:"status"`
    Priority    string    `json:"priority"`
    DueDate     time.Time `json:"due_date,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    CompletedAt time.Time `json:"completed_at,omitempty"`
}

// TodoInput represents the input for creating a new todo
type TodoInput struct {
    Task     string    `json:"task" validate:"required,min=3,max=100"`
    Priority string    `json:"priority" validate:"oneof=low medium high"`
    DueDate  time.Time `json:"due_date,omitempty"`
}

// StatusUpdate represents a status update request
type StatusUpdate struct {
    Item   string `json:"item" validate:"required"`
    Status string `json:"status" validate:"required,oneof=TO_BE_STARTED IN_PROGRESS COMPLETED"`
}

// Constants for validation
const (
    StatusToBeStarted = "TO_BE_STARTED"
    StatusInProgress  = "IN_PROGRESS"
    StatusCompleted   = "COMPLETED"

    PriorityLow    = "low"
    PriorityMedium = "medium"
    PriorityHigh   = "high"
)

// ValidStatus checks if a status is valid
func ValidStatus(status string) bool {
    switch status {
    case StatusToBeStarted, StatusInProgress, StatusCompleted:
        return true
    default:
        return false
    }
}

// ValidPriority checks if a priority is valid
func ValidPriority(priority string) bool {
    switch priority {
    case PriorityLow, PriorityMedium, PriorityHigh:
        return true
    default:
        return false
    }
}