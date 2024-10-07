package router

import (
	"time-tracker-app/internal/transport/http/handler"

	"github.com/gofiber/fiber/v3"
)

type UserRouter struct {
	userHandler *handler.UserHandler
}

func NewUserRouter(userHandler *handler.UserHandler) *UserRouter {
	return &UserRouter{userHandler}
}

func (ur *UserRouter) setUserRoutes(router fiber.Router) {
	users := router.Group("/users")
	{
		users.Post("/", ur.userHandler.CreateUser)
		users.Get("/", ur.userHandler.GetAllUsers)
		users.Get("/:userID", ur.userHandler.GetUserByID)
		// users.Delete("/:userID", ur.userHandler.CreateUser)
	}
}
