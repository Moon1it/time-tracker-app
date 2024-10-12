package handler

import (
	"context"
	"time-tracker-app/internal/converter"
	"time-tracker-app/internal/domain"
	"time-tracker-app/internal/transport/http/dto"

	"github.com/gofiber/fiber/v3"
)

type ITaskService interface {
	CreateUser(c context.Context, info *domain.TaskInfo) (string, error)
	GetTasksByUserID(c context.Context, userID string) ([]domain.Task, error)
}

type TaskHandler struct {
	taskService ITaskService
}

func NewTaskHandler(userService ITaskService) *TaskHandler {
	return &TaskHandler{userService}
}

func (th *TaskHandler) CreateTask(c fiber.Ctx) error {
	payload := new(dto.TaskInfo)
	if err := c.Bind().Body(payload); err != nil {
		return err
	}

	userID, err := th.taskService.CreateUser(c.Context(), converter.ToTaskInfoFromHandler(payload))
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"": ""})
	}

	return c.Status(fiber.StatusOK).JSON(userID)
}

// func (th *TaskHandler) GetTasksByUserID(c fiber.Ctx) error {
// 	userID := c.Params("userID")

// 	tasks, err := th.taskService.GetTasksByUserID(c.Context(), userID)
// 	if err != nil {
// 	}

// 	return c.Status(fiber.StatusOK).JSON(converter.ToGetAllTasksResponseFromService(tasks))
// }

// func (uh *TaskHandler) GetTaskByID(c fiber.Ctx) error {
// 	userID := c.Params("userID")

// 	user, err := uh.userService.GetTaskByID(c.Context(), userID)
// 	if err != nil {
// 	}

// 	return c.Status(fiber.StatusOK).JSON(user)
// }
