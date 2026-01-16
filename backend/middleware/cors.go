package middleware

import (
	"net/http"
	"os"

	"github.com/uptrace/bunrouter"
)

func CORS() bunrouter.MiddlewareFunc {
	return func(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
		return func(w http.ResponseWriter, req bunrouter.Request) error {
			origin := req.Header.Get("Origin")

			var allowedOrigins []string
			if os.Getenv("APP_ENV") == "dev" {
				allowedOrigins = []string{
					"http://localhost:3000",
					"http://localhost:5173",
					"http://localhost:5174",
				}
			} else {
				allowedOrigins = []string{os.Getenv("FRONTEND_URL")}
			}
			isAllowed := false
			for _, allowed := range allowedOrigins {
				if origin == allowed {
					isAllowed = true
					break
				}
			}

			if isAllowed {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
				w.Header().Set("Access-Control-Max-Age", "3600")
			}

			if req.Method == "OPTIONS" {
				w.WriteHeader(http.StatusNoContent)
				return nil
			}
			return next(w, req)
		}
	}
}
