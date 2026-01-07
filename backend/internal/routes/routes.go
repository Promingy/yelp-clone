package routes

import (
	"github.com/promingy/yelp-clone/backend/internal/handlers"

	"github.com/uptrace/bunrouter"
)

func SetupRoutes(router *bunrouter.Router) {
	router.WithGroup("/api", func(group *bunrouter.Group) {
		// /api/users
		group.WithGroup("/users", func(group *bunrouter.Group) {
			// db, rowLimit, rateLimit
			userHandler := &handlers.UserHandler{}

			group.POST("/", userHandler.CreateUserHandler)
			group.GET("/", userHandler.ShowUserHandler)
			group.GET("/list", userHandler.ListUsersHandler)
			group.PUT("/", userHandler.UpdateUserHandler)
			group.DELETE("/", userHandler.DeleteUserHandler)
		})
	})
}
