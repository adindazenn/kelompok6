// model/user.go

package model

import "time"

type User struct {
    ID        uint       `gorm:"primary_key" json:"id"`
    FullName  string     `json:"full_name" binding:"required"`
    Email     string     `json:"email" binding:"required,email" gorm:"unique_index"`
    Password  string     `json:"password" binding:"required,min=6"`
    Role      string     `json:"role" binding:"required,oneof=admin member"`
    CreatedAt time.Time  `json:"created_at"`
    UpdatedAt time.Time  `json:"updated_at"`
    Tasks     []Task     `json:"tasks"`
}

type RegisterRequest struct {
    FullName string `json:"full_name"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type LoginRequest struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required"`
}

type UpdateAccountRequest struct {
    FullName string `json:"full_name"`
    Email    string `json:"email" binding:"email"`
}

type UserResponse struct {
    ID        uint       `json:"id"`
    FullName  string    `json:"full_name"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"created_at"`
}

type UpdateUserResponse struct {
    ID        uint       `json:"id"`
    FullName  string    `json:"full_name"`
    Email     string    `json:"email"`
    UpdatedAt time.Time `json:"updated_at"`
}

