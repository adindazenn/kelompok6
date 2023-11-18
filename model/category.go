package model

import (
	"time"
)

type Category struct {
	ID              uint      `gorm:"primary_key" json:"id"`
	Type            string    `json:"type" binding:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Tasks []Task `json:"tasks"`
}

type CategoryResponse struct {
    ID        uint       `json:"id"`
    Type      string     `json:"type"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    Tasks     []TaskInfo `json:"Tasks"`
}

type CreateCategoryResponse struct {
    ID        uint       `json:"id"`
    Type      string     `json:"type"`
    CreatedAt time.Time  `json:"created_at"`
}

type UpdateCategoryResponse struct {
    ID        uint       `json:"id"`
    Type      string     `json:"type"`
    UpdatedAt time.Time  `json:"updated_at"`
}



