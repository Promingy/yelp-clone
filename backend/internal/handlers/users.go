package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/promingy/yelp-clone/backend/internal/models"
	v "github.com/promingy/yelp-clone/backend/internal/validation"
	"github.com/uptrace/bun"
	"github.com/uptrace/bunrouter"
)

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
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Bio         string `json:"bio"`
	Country     string `json:"country"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	ProfilePic  string `json:"profile_pic"`
}


// / ----------- START POST ROUTE HANDLERS ---------
func (h *UserHandler) CreateNewUser(w http.ResponseWriter, req bunrouter.Request) error {
	var input UserInput

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		return err
	}
	defer req.Body.Close()

	user := &models.User{
		Email:       input.Email,
		Password:    input.Password,
	}
	if errs := v.Validate(user); len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, map[string]map[string]string{
			"errors": errs,
		})
	}
	_, err := h.db.NewInsert().Model(user).Exec(req.Context())
	if err != nil {
		return bunrouter.JSON(w, map[string]string{"error": err.Error()})
	}

	profile := &models.Profile{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Bio:         input.Bio,
		Country:     input.Country,
		City:        input.City,
		PhoneNumber: input.PhoneNumber,
		State:       input.State,
		ZipCode:     input.ZipCode,
		ProfilePic:  input.ProfilePic,
	}
	if errs := v.Validate(profile); len(errs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, map[string]map[string]string{"errors": errs,})
	}
	_, err = h.db.NewInsert().Model(profile).Exec(req.Context())
	if err != nil {
		return bunrouter.JSON(w, map[string]string{"error": err.Error()})
	}

	return bunrouter.JSON(w, profile)
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
