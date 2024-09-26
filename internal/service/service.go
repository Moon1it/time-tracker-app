package service

import (
	"time-tracker-app/internal/repository"
)

type Service struct {
	UserService *UserService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repository.UserRepository),
	}
}
