package repository

import (
	"context"
	db "time-tracker-app/internal/db/sqlc"
	"time-tracker-app/internal/domain"
	"time-tracker-app/internal/repository/converter"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserRepository struct {
	querier db.Querier
}

func NewUserRepository(querier db.Querier) *UserRepository {
	return &UserRepository{querier}
}

func (ur *UserRepository) CreateUser(c context.Context, userID uuid.UUID, userInfo *domain.UserInfo) error {
	arg := db.CreateUserParams{
		Uuid:      pgtype.UUID{Bytes: userID, Valid: true},
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
	}

	err := ur.querier.CreateUser(c, arg)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) GetUsers(c context.Context) ([]domain.User, error) {
	users, err := ur.querier.GetUsers(c)
	if err != nil {
		return nil, err
	}

	return converter.ToUsersFromRepo(users), nil
}
