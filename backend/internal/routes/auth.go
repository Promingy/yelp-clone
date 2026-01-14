package routes

import (
	"github.com/promingy/yelp-clone/backend/internal/handlers"
	"github.com/promingy/yelp-clone/backend/middleware"
	"github.com/uptrace/bunrouter"
)

func RegisterAuthRoutes(api *bunrouter.Group, h *handlers.AuthHandler, m *middleware.AuthMiddleware) {
	api.WithGroup("/auth", func(auth *bunrouter.Group) {
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/refresh", h.RefreshToken)

		auth.GET("/auth/me", m.RequireAuth(h.GetCurrentUser))
	})
}
