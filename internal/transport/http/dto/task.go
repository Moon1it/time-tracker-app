package dto

import "time"

type TaskInfo struct {
	UserID string `json:"userID" validate:"required"`
	Name   string `json:"name" validate:"required"`
}

type TaskResponse struct {
	ID          string     `json:"id" validate:"required"`
	Info        TaskInfo   `json:"info" validate:"required"`
	Duration    int        `json:"duration" validate:"required"`
	IsComplited bool       `json:"isComplited" validate:"required"`
	CreatedAt   time.Time  `json:"createdAt" validate:"required"`
	UpdatedAt   *time.Time `json:"updatedAt" validate:"required"`
}

type GetTasksByUserIDResponse struct {
	Size  int            `json:"size" validate:"required"`
	Tasks []TaskResponse `json:"tasks" validate:"required"`
}
