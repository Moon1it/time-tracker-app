package service

import (
	"context"
	"time-tracker-app/internal/domain"
	"time-tracker-app/internal/repository"

	"github.com/google/uuid"
)

// type IUserRepository interface {
// 	CreateUser(ctx context.Context, arg repository.CreateUserParams) error
// 	GetUserByUUID(ctx context.Context, userUuid pgtype.UUID) (repository.User, error)
// 	GetUsers(ctx context.Context) (repository.User, error)
// }

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) CreateUser(ctx context.Context, info *domain.UserInfo) (string, error) {
	userID := uuid.New()

	err := us.userRepository.CreateUser(ctx, userID, info)
	if err != nil {
		return "", err
	}

	return userID.String(), nil
}

func (us *UserService) GetAllUsers(c context.Context) ([]domain.User, error) {
	users, err := us.userRepository.GetUsers(c)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// func (us *UserService) GetUserByID(ctx context.Context, userID string) (string, error) {
// 	user, err := us.userRepository.GetUserByUUID(ctx, converter.ToPgUUIDFromService(userID))
// 	if err != nil {
// 		return "", err
// 	}

// 	return user, nil
// }
