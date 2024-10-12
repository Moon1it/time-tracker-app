package converter

import (
	db "time-tracker-app/internal/db/sqlc"
	"time-tracker-app/internal/domain"

	"github.com/google/uuid"
)

func ToUsersFromRepo(dbUsers []db.User) []domain.User {
	users := make([]domain.User, 0, len(dbUsers))

	for _, user := range dbUsers {
		users = append(users, domain.User{
			ID: uuid.UUID(user.Uuid.Bytes).String(),
			Info: domain.UserInfo{
				FirstName: user.FirstName,
				LastName:  user.LastName,
			},
			CreatedAt: user.CreatedAt.Time,
			UpdatedAt: &user.UpdatedAt.Time,
		})
	}

	return users
}

func ToUserFromRepo(user *db.User) *domain.User {
	return &domain.User{
		ID: uuid.UUID(user.Uuid.Bytes).String(),
		Info: domain.UserInfo{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: &user.UpdatedAt.Time,
	}
}
