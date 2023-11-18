package model

import "time"

type Task struct {
    ID          uint      `gorm:"primary_key" json:"id"`
    Title       string    `json:"title" binding:"required"`
    Description string    `json:"description" binding:"required"`
    Status      bool      `json:"status"`
    UserID      uint      `json:"user_id"`
    CategoryID  uint      `json:"category_id"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateTaskRequest struct {
    Title       string `json:"title" binding:"required"`
    Description string `json:"description" binding:"required"`
}

// UpdateTaskCategoryRequest adalah model untuk request update category task.
type UpdateTaskCategoryRequest struct {
    CategoryID uint `json:"category_id" binding:"required"`
}

// UpdateTaskStatusRequest adalah model untuk request update status task.
type UpdateTaskStatusRequest struct {
    Status bool `json:"status" binding:"required"`
}

type TaskInfo struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    UserID      uint      `json:"user_id"`
    CategoryID  uint      `json:"category_id"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTaskResponse struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    Status      bool      `json:"status"`
    Description string    `json:"description"`
    UserID      uint      `json:"user_id"`
    CategoryID  uint      `json:"category_id"`
    CreatedAt   time.Time `json:"created_at"`
}

type TaskResponse struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    Status      bool      `json:"status"`
    Description string    `json:"description"`
    UserID      uint      `json:"user_id"`
    CategoryID  uint      `json:"category_id"`
    CreatedAt   time.Time `json:"created_at"`
    User        struct {
        ID       uint   `json:"id"`
        Email    string `json:"email"`
        FullName string `json:"full_name"`
    } `json:"User"`
}

type TaskUpdateResponse struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Status      bool      `json:"status"`
    UserID      uint      `json:"user_id"`
    CategoryID  uint      `json:"category_id"`
    UpdatedAt   time.Time `json:"updated_at"`
}
