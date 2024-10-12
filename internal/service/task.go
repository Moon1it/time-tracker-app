package service

import (
	"context"
	"time-tracker-app/internal/domain"
)

type ITaskRepository interface {
	CreateTask(c context.Context, TaskInfo *domain.TaskInfo) (string, error)
	GetTasksByUserID(c context.Context, userID string) ([]domain.Task, error)
}

type TaskService struct {
	taskRepository ITaskRepository
}

func NewTaskService(taskRepository ITaskRepository) *TaskService {
	return &TaskService{taskRepository}
}

func (us *TaskService) CreateUser(c context.Context, info *domain.TaskInfo) (string, error) {
	taskID, err := us.taskRepository.CreateTask(c, info)
	if err != nil {
		return "", err
	}

	return taskID, nil
}

func (ts *TaskService) GetTasksByUserID(c context.Context, userID string) ([]domain.Task, error) {
	users, err := ts.taskRepository.GetTasksByUserID(c, userID)
	if err != nil {
		return nil, err
	}

	return users, nil
}
