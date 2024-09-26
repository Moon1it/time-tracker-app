package handler

import (
	"time-tracker-app/internal/converter"
	"time-tracker-app/internal/service"
	"time-tracker-app/internal/transport/http/dto"

	"github.com/gofiber/fiber/v3"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (uh *UserHandler) CreateUser(c fiber.Ctx) error {
	payload := new(dto.CreateUserPaylaod)
	if err := c.Bind().Body(payload); err != nil {
		return err
	}

	userID, err := uh.userService.CreateUser(c.Context(), converter.ToUserInfoFromHandler(payload))
	if err != nil {
		return err
	}

	return c.Status(200).JSON(userID)
}

func (uh *UserHandler) GetAllUsers(c fiber.Ctx) error {
	users, err := uh.userService.GetAllUsers(c.Context())
	if err != nil {
		return nil
	}

	return c.Status(200).JSON(converter.ToGetAllUsersResponseFromService(users))
}

// func (uh *UserHandler) GetUserByID(c fiber.Ctx) error {
// userIDRaw := c.Params("userID")
// userID, err := uuid.Parse(userIDRaw)
// if err != nil {
// 	c.Status(500).JSON("...")
// }

// user, ok := userMap[userID]
// if !ok {
// 	return c.Status(500).JSON(map[string]string{"message": "user not found"})
// }

// 	// return c.Status(200).JSON(user)
// }
