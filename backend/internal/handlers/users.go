package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/promingy/yelp-clone/backend/internal/services"
	"github.com/uptrace/bunrouter"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService}
}


// #endregion

type CreateUserRequest struct {
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
	var input CreateUserRequest

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, map[string]string{"error": "Invalid request body"})
	}
	defer req.Body.Close()
	serviceInput := services.CreateUserInput{
		FirstName:   input.FirstName,
		LastName:    input.LastName,
		Email:       input.Email,
		Password:    input.Password,
		PhoneNumber: input.PhoneNumber,
		Bio:         input.Bio,
		Country:     input.Country,
		City:        input.City,
		State:       input.State,
		ZipCode:     input.ZipCode,
		ProfilePic:  input.ProfilePic,
	}

	result, err := h.userService.CreateUser(req.Context(), serviceInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, map[string]string{"error": err.Error()})
	}

	return bunrouter.JSON(w, map[string]interface{}{
		"user": map[string]interface{}{
			"id": result.User.ID,
			"email": result.User.Email,
		},
		"profile": result.Profile,
	})
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
