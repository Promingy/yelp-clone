package main

import (
	"log"
	"net/http"

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

	/// Initialize Handlers
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	/// Setup router
	router := bunrouter.New()
	router.WithGroup("/api", func(api *bunrouter.Group) {
		routes.RegisterUserRoutes(api, userHandler)
		routes.RegisterAuthRoutes(api, authHandler, authMiddleware)
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}
