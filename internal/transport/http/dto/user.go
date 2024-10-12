package dto

import "time"

type UserInfo struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type UserResponse struct {
	ID        string     `json:"id" validate:"required"`
	Info      UserInfo   `json:"info" validate:"required"`
	CreatedAt time.Time  `json:"createdAt" validate:"required"`
	UpdatedAt *time.Time `json:"updatedAt" validate:"required"`
}

type GetAllUsersResponse struct {
	Size  int            `json:"size" validate:"required"`
	Users []UserResponse `json:"users" validate:"required"`
}
