package routes

import (
	"github.com/promingy/yelp-clone/backend/internal/handlers"
	"github.com/uptrace/bunrouter"
)

func RegisterUserRoutes(api *bunrouter.Group, h *handlers.UserHandler) {
	api.WithGroup("/users", func(users *bunrouter.Group) {
		users.GET("/list", h.ListUsersHandler)
		users.POST("/new/", h.CreateNewUser)
		users.GET("/", h.ShowUserHandler)
		users.PUT("/", h.UpdateUserHandler)
		users.DELETE("/", h.DeleteUserHandler)
	})
}
