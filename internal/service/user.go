package service

import (
	"context"
	"time-tracker-app/internal/domain"
)

type IUserRepository interface {
	CreateUser(c context.Context, userInfo *domain.UserInfo) (string, error)
	GetUsers(c context.Context) ([]domain.User, error)
	GetUserByID(c context.Context, userID string) (*domain.User, error)
}

type UserService struct {
	userRepository IUserRepository
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{userRepository}
}

func (us *UserService) CreateUser(c context.Context, info *domain.UserInfo) (string, error) {
	userID, err := us.userRepository.CreateUser(c, info)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (us *UserService) GetAllUsers(c context.Context) ([]domain.User, error) {
	users, err := us.userRepository.GetUsers(c)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	user, err := us.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
