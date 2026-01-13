package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/promingy/yelp-clone/backend/internal/db"
	"github.com/promingy/yelp-clone/backend/internal/handlers"
	"github.com/promingy/yelp-clone/backend/internal/repositories"
	"github.com/promingy/yelp-clone/backend/internal/routes"
	"github.com/promingy/yelp-clone/backend/internal/services"

	"github.com/uptrace/bunrouter"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	validator := validator.New(validator.WithRequiredStructEnabled())

	/// Initialize Repositories
	userRepo := repositories.NewUserRepository(db)
	profileRepo := repositories.NewProfileRepository(db)
	
	/// Initialize Services
	userService := services.NewUserService(userRepo, profileRepo, validator)
	
	/// Initialize Handlers
	userHandler := handlers.NewUserHandler(userService)
	
	/// Setup router
	router := bunrouter.New()
	routes.SetupRoutes(router, routes.Handlers{
		Users: userHandler,
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)

	// example
	// router.GET("/users/:id", handlers.ShowUserHandler)
	// can be retrieved with
	// params := req.Params()
	// path := params.ByName("<param_name>") for *param
	// id, err := params.Int64("id")
}
