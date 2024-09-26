package handler

import "time-tracker-app/internal/service"

type Handler struct {
	UserHandler *UserHandler
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		UserHandler: NewUserHandler(services.UserService),
	}
}
