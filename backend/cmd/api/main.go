package main

import (
	"log"
	"net/http"

	"github.com/promingy/yelp-clone/backend/internal/db"
	"github.com/promingy/yelp-clone/backend/internal/routes"

	"github.com/uptrace/bunrouter"
)

func main() {
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := bunrouter.New()

	routes.SetupRoutes(router)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)

	// example
	// router.GET("/users/:id", handlers.ShowUserHandler)
	// can be retrieved with
	// params := req.Params()
	// path := params.ByName("<param_name>") for *param
	// id, err := params.Int64("id")
}
