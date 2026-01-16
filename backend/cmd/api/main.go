package main

import (
	"log"
	"net/http"
	"os"

	"filippo.io/csrf"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/promingy/yelp-clone/backend/config"
	"github.com/promingy/yelp-clone/backend/internal/db"
	"github.com/promingy/yelp-clone/backend/internal/handlers"
	"github.com/promingy/yelp-clone/backend/internal/repositories"
	"github.com/promingy/yelp-clone/backend/internal/routes"
	"github.com/promingy/yelp-clone/backend/internal/services"
	"github.com/promingy/yelp-clone/backend/middleware"

	"github.com/uptrace/bunrouter"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	jwtConfig := config.LoadJWTConfig()
	validator := validator.New(validator.WithRequiredStructEnabled())

	/// Initialize Repositories
	userRepo := repositories.NewUserRepository(db)
	profileRepo := repositories.NewProfileRepository(db)

	/// Initialize Services
	userService := services.NewUserService(userRepo, profileRepo, validator)
	authService := services.NewAuthService(userRepo, userService, jwtConfig)

	/// Middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)
	corsMiddleWare := middleware.CORS()

	/// Initialize Handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	/// Setup router
	router := bunrouter.New()
	router.Use(corsMiddleWare)

	router.WithGroup("/api", func(api *bunrouter.Group) {
		routes.RegisterUserRoutes(api, userHandler)
		routes.RegisterAuthRoutes(api, authHandler, authMiddleware)
	})

	protection := csrf.New()

	if os.Getenv("APP_ENV") == "dev" {
		protection.AddTrustedOrigin("http://localhost:3000")
		protection.AddTrustedOrigin("http://localhost:5173")
		protection.AddTrustedOrigin("http://localhost:5174")
	} else {
		if frontendUrl := os.Getenv("FRONTEND_URL"); frontendUrl != "" {
			protection.AddTrustedOrigin(frontendUrl)
		}
	}

	handler := protection.Handler(router)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", handler)
}
