package backend

import (
	"fmt"
	"net/http"

	"github.com/promingy/yelp-clone/backend/internal/handlers"

	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func main() {
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
	)

	router.GET("/", func(w http.ResponseWriter, req bunrouter.Request) error {
		// req embeds *http.Request and has all the same fields and methods

		fmt.Println(req.Method, req.Route(), req.Params().Map())
		return nil
	})

	router.POST("/users", handlers.CreateUserHandler)
	router.GET("/users", handlers.ShowUserHandler)
	router.PUT("/users", handlers.UpdateUserHandler)
	router.DELETE("/users", handlers.DeleteUserHandler)
}
