package router

import (
	"time-tracker-app/internal/transport/http/handler"

	"github.com/gofiber/fiber/v3"
)

type Router struct {
	UserRouter *UserRouter
}

func NewRouter(handler *handler.Handler) *Router {
	return &Router{
		UserRouter: NewUserRouter(handler.UserHandler),
	}
}

func (r *Router) SetRoutes(router *fiber.App) {
	api := router.Group("/api")
	r.UserRouter.setUserRoutes(api)
}
