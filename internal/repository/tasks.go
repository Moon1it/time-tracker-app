package repository

import (
	"context"
	db "time-tracker-app/internal/db/sqlc"
	"time-tracker-app/internal/domain"
	"time-tracker-app/internal/repository/converter"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskRepository struct {
	querier db.Querier
}

func NewTaskRepository(querier db.Querier) *TaskRepository {
	return &TaskRepository{querier}
}

func (tr *TaskRepository) CreateTask(c context.Context, TaskInfo *domain.TaskInfo) (string, error) {
	userUUID, err := uuid.Parse(TaskInfo.UserID)
	if err != nil {
		return "", err
	}

	arg := db.CreateTaskParams{
		UserUuid: pgtype.UUID{Bytes: userUUID, Valid: true},
		Name:     TaskInfo.Name,
	}
	taskPgUUID, err := tr.querier.CreateTask(c, arg)
	if err != nil {
		return "", err
	}

	return uuid.UUID(taskPgUUID.Bytes).String(), nil
}

func (tr *TaskRepository) GetTasksByUserID(c context.Context, userID string) ([]domain.Task, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	tasks, err := tr.querier.GetTasksByUserUUID(c, pgtype.UUID{Bytes: userUUID, Valid: true})
	if err != nil {
		return nil, err
	}

	return converter.ToTasksFromRepo(tasks), nil
}
