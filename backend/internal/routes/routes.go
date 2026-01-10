package routes

import (
	"github.com/promingy/yelp-clone/backend/internal/handlers"
	"github.com/uptrace/bunrouter"
)

type Handlers struct {
	Users *handlers.UserHandler
}

func SetupRoutes(router *bunrouter.Router, h Handlers) {
	router.WithGroup("/api", func(api *bunrouter.Group) {
		RegisterUserRoutes(api, h.Users)
	})
}
