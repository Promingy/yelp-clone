package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/promingy/yelp-clone/backend/internal/services"
	"github.com/uptrace/bunrouter"
)

type AuthMiddleware struct {
	authService *services.AuthService
}

func NewAuthMiddleware (authService *services.AuthService) *AuthMiddleware {
	return &AuthMiddleware{authService}
}

func (m *AuthMiddleware) RequireAuth(next bunrouter.HandlerFunc) bunrouter.HandlerFunc{
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return bunrouter.JSON(w, map[string]string{"error": "Missing authorization header"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			return bunrouter.JSON(w, map[string]string{"error": "Invalid authorization header format"})
		}

		tokenString := parts[1]

		claims, err := m.authService.ValidateAccessToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return bunrouter.JSON(w, map[string]string{"error": "Invalid or expired token"})
		}

		ctx := context.WithValue(req.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "email", claims.Email)

		req = req.WithContext(ctx)

		return next(w, req)
	}
}