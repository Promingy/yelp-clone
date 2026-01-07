package handlers

import (
	"net/http"

	"github.com/uptrace/bun"
	"github.com/uptrace/bunrouter"
)

type UserHandler struct {
	db        *bun.DB
	rowLimit  int
	rateLimit int
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}

func (h *UserHandler) ShowUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}

func (h *UserHandler) ListUsersHandler(w http.ResponseWriter, req bunrouter.Request) error {
	type UserResponse struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	}

	data := UserResponse{
		ID: 1,
		Name: "John Doe", 
		Active: true,
	}
	
	return bunrouter.JSON(w, data)
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, req bunrouter.Request) error {
	return nil
}
