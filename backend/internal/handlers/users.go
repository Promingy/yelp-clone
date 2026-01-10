package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "github.com/promingy/yelp-clone/backend/internal/models"
	"github.com/promingy/yelp-clone/backend/internal/models"
	"github.com/uptrace/bun"
	"github.com/uptrace/bunrouter"
)

// #region
type UserHandler struct {
	db        *bun.DB
	rowLimit  int
	rateLimit int
}

func NewUserHandler(db *bun.DB) *UserHandler {
	return &UserHandler{
		db:       db,
		rowLimit: 100,
	}
}

// #endregion

type UserInput struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Bio        string `json:"bio"`
	Country    string `json:"country"`
	State      string `json:"state"`
	ZipCode    string `json:"zipcode"`
	ProfilePic string `json:"profile_pic"`
}

// / ----------- START POST ROUTE HANDLERS ---------
func (h *UserHandler) CreateNewUser(w http.ResponseWriter, req bunrouter.Request) error {
	var input UserInput

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		return err
	}
	defer req.Body.Close()

	user := &models.User{FirstName: input.FirstName}

	if errs := user.Validate(); len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, map[string]map[string]string{
			"errors": errs,
		})
	}

	_, err := h.db.NewInsert().Model(user).Exec(req.Context())
	if err != nil {
		panic(err)
	}

	return bunrouter.JSON(w, req.Body)
}

/// ----------- END POST ROUTE HANDLERS ---------

/// ----------- START GET ROUTE HANDLERS ---------

func (h *UserHandler) ShowUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	fmt.Print("ShowUserHandler")
	return nil
}

func (h *UserHandler) ListUsersHandler(w http.ResponseWriter, req bunrouter.Request) error {
	type UserResponse struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}

	data := UserResponse{
		ID:     1,
		Name:   "John Doe",
		Active: true,
	}

	return bunrouter.JSON(w, data)
}

/// ----------- END GET ROUTE HANDLERS ---------

/// ----------- START UPDATE ROUTE HANDLERS ---------

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	fmt.Print("UpdateUserHandler")
	return nil
}

/// ----------- END UPDATE ROUTE HANDLERS ---------

/// ----------- START DELETE ROUTE HANDLERS ---------

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	fmt.Print("DeleteUserHandler")
	return nil
}

/// ----------- END DELETE ROUTE HANDLERS ---------
