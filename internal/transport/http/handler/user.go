package handler

import (
	"context"
	"fmt"
	"time-tracker-app/internal/converter"
	"time-tracker-app/internal/domain"
	"time-tracker-app/internal/transport/http/dto"

	"github.com/gofiber/fiber/v3"
)

type IUserService interface {
	CreateUser(c context.Context, info *domain.UserInfo) (string, error)
	GetAllUsers(c context.Context) ([]domain.User, error)
	GetUserByID(ctx context.Context, userID string) (*domain.User, error)
}

type UserHandler struct {
	userService IUserService
}

func NewUserHandler(userService IUserService) *UserHandler {
	return &UserHandler{userService}
}

func (uh *UserHandler) CreateUser(c fiber.Ctx) error {
	payload := new(dto.CreateUserPaylaod)
	if err := c.Bind().Body(payload); err != nil {
		return err
	}

	fmt.Println(payload)

	userID, err := uh.userService.CreateUser(c.Context(), converter.ToUserInfoFromHandler(payload))
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"": ""})
	}

	return c.Status(fiber.StatusOK).JSON(userID)
}

func (uh *UserHandler) GetAllUsers(c fiber.Ctx) error {
	users, err := uh.userService.GetAllUsers(c.Context())
	if err != nil {
	}

	return c.Status(fiber.StatusOK).JSON(converter.ToGetAllUsersResponseFromService(users))
}

func (uh *UserHandler) GetUserByID(c fiber.Ctx) error {
	userID := c.Params("userID")

	user, err := uh.userService.GetUserByID(c.Context(), userID)
	if err != nil {
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
