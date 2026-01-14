package handlers

import (
	"encoding/json"
	"net/http"

	e "github.com/promingy/yelp-clone/backend/internal/errors"
	"github.com/promingy/yelp-clone/backend/internal/services"
	"github.com/uptrace/bunrouter"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, req bunrouter.Request) error {
	var input LoginRequest

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, map[string]string{"error": "Invalid request body"})
	}
	defer req.Body.Close()

	result, err := h.authService.Login(req.Context(), services.LoginInput{
		Email: input.Email,
		Password: input.Password,
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, map[string]string{"error": err.Error()})
	}

	return bunrouter.JSON(w, result)
}

func (h *AuthHandler) Register(w http.ResponseWriter, req bunrouter.Request) error {
	var input services.CreateUserInput

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, map[string]string{"error": "Invalid request body"})
	}
	defer req.Body.Close()

	result, err := h.authService.Register(req.Context(), input)
	if err != nil {
		if validationErr, ok := err.(*e.ValidationError); ok {
			w.WriteHeader(http.StatusBadRequest)
			return bunrouter.JSON(w, map[string]map[string]string{
				"errors": validationErr.Errors,
			})
		}

		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, map[string]string{"error": err.Error()})
	}

	return bunrouter.JSON(w, result)
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, req bunrouter.Request) error {
	var input RefreshTokenRequest

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, map[string]string{"error": "Invalid request body"})
	}
	defer req.Body.Close()

	result, err := h.authService.RefreshToken(req.Context(), input.RefreshToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, map[string]string{"error": err.Error()})
	}

	return bunrouter.JSON(w, result)
}

func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, req bunrouter.Request) error {
	userID := req.Context().Value("user_id").(int64)

	user, err := h.authService.UserRepo.FindById(req.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return bunrouter.JSON(w, map[string]string{"error": "User not found"})
	}
	
	profile, err := h.authService.UserRepo.GetProfileByUserId(req.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return bunrouter.JSON(w, map[string]string{"error": "User not found"})
	}

	return bunrouter.JSON(w, map[string]interface{}{
		"User": services.UserResponse{
			ID: user.ID,
			Email: user.Email,
			FirstName: profile.FirstName,
			LastName: profile.LastName,
			ProfilePic: profile.ProfilePic,
		},
	})
}
