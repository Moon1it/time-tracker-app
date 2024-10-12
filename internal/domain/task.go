package domain

import "time"

type Task struct {
	ID          string
	Info        TaskInfo
	Duration    int
	IsComplited bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type TaskInfo struct {
	UserID string
	Name   string
}
