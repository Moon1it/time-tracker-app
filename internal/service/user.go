package service

import (
	"context"
	"time-tracker-app/internal/domain"

	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(c context.Context, userID uuid.UUID, userInfo *domain.UserInfo) error
	GetUsers(c context.Context) ([]domain.User, error)
	GetUserByUUID(c context.Context, userUUID uuid.UUID) (*domain.User, error)
}

type UserService struct {
	userRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) CreateUser(c context.Context, info *domain.UserInfo) (string, error) {
	userUUID := uuid.New()

	err := us.userRepository.CreateUser(c, userUUID, info)
	if err != nil {
		return "", err
	}

	return userUUID.String(), nil
}

func (us *UserService) GetAllUsers(c context.Context) ([]domain.User, error) {
	users, err := us.userRepository.GetUsers(c)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}

	user, err := us.userRepository.GetUserByUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
