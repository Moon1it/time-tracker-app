package converter

import (
	db "time-tracker-app/internal/db/sqlc"
	"time-tracker-app/internal/domain"

	"github.com/google/uuid"
)

func ToTasksFromRepo(dbTasks []db.Task) []domain.Task {
	tasks := make([]domain.Task, 0, len(dbTasks))

	for _, task := range dbTasks {
		tasks = append(tasks, domain.Task{
			ID: uuid.UUID(task.Uuid.Bytes).String(),
			Info: domain.TaskInfo{
				UserID: uuid.UUID(task.UserUuid.Bytes).String(),
				Name:   task.Name,
			},
			Duration:    int(task.Duration),
			IsComplited: task.IsCompleted,
			CreatedAt:   task.CreatedAt.Time,
			UpdatedAt:   &task.UpdatedAt.Time,
		})
	}

	return tasks
}

// func ToTaskFromRepo(user *db.User) *domain.User {
// return &domain.User{
// 	ID: uuid.UUID(user.Uuid.Bytes).String(),
// 	Info: domain.UserInfo{
// 		FirstName: user.FirstName,
// 		LastName:  user.LastName,
// 	},
// 	CreatedAt: user.CreatedAt.Time,
// 	UpdatedAt: &user.UpdatedAt.Time,
// }
// }
