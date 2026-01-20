package handlers

import (
	"encoding/json"
	"net/http"

	e "github.com/promingy/yelp-clone/backend/internal/errors"
	"github.com/promingy/yelp-clone/backend/internal/services"
	"github.com/promingy/yelp-clone/backend/internal/utils"
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
		return bunrouter.JSON(w, e.NewSingleError("Invalid request body"))
	}
	defer req.Body.Close()

	result, err := h.authService.Login(req.Context(), services.LoginInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, e.NewSingleError(err.Error()))
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    result.AccessToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   15 * 60,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    result.RefreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/api/auth",
		MaxAge:   7 * 24 * 60 * 60,
	})

	return bunrouter.JSON(w, result.User)
}

func (h *AuthHandler) Register(w http.ResponseWriter, req bunrouter.Request) error {
	var input services.CreateUserInput

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, e.NewSingleError("Invalid request body"))
	}
	defer req.Body.Close()

	result, err := h.authService.Register(req.Context(), input)
	if err != nil {
		if validationErr, ok := err.(*e.ValidationError); ok {
			w.WriteHeader(http.StatusBadRequest)
			return bunrouter.JSON(w, e.NewMultiError(validationErr.Errors))
		}

		w.WriteHeader(http.StatusBadRequest)
		return bunrouter.JSON(w, e.NewSingleError(err.Error()))
	}

	return bunrouter.JSON(w, result.User)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, req bunrouter.Request) error {
	h.authService.Logout(req.Context())

	http.SetCookie(w, &http.Cookie{
		Name:   "access_token",
		Path:   "/",
		MaxAge: -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:   "refresh_token",
		Path:   "/api/auth",
		MaxAge: -1,
	})

	return bunrouter.JSON(w, map[string]string{"message": "Successfully logged out"})
}

func (h *AuthHandler) DeleteCurrentUser(w http.ResponseWriter, req bunrouter.Request) error {
	userID, ok := utils.GetUserID(req.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, e.NewSingleError("Unauthorized"))
	}

	err := h.authService.DeleteCurrentUser(req.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, e.NewSingleError("Failed to delete user"))
	}

	return bunrouter.JSON(w, map[string]string {
		"message": "User successfully deleted",
	})
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, req bunrouter.Request) error {
	cookie, err := req.Cookie("refresh_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, e.NewSingleError("No refresh token"))
	}

	result, err := h.authService.RefreshToken(req.Context(), cookie.Value)
	if err != nil {
		http.SetCookie(w, &http.Cookie{
			Name:   "access_token",
			MaxAge: -1,
		})
		http.SetCookie(w, &http.Cookie{
			Name:   "refresh_token",
			MaxAge: -1,
		})

		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, e.NewSingleError("Unauthorized"))
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    result.AccessToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
		MaxAge:   15 * 60,
	})

	return bunrouter.JSON(w, result.User)
}

func (h *AuthHandler) GetCurrentUser(w http.ResponseWriter, req bunrouter.Request) error {
	userID, ok := utils.GetUserID(req.Context())
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return bunrouter.JSON(w, e.NewSingleError("Unauthorized"))
	}

	user, err := h.authService.GetCurrentUser(req.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return bunrouter.JSON(w, e.NewSingleError(err.Error()))
	}

	return bunrouter.JSON(w, user)
}
