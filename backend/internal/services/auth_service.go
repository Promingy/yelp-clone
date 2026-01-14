package services

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/promingy/yelp-clone/backend/config"
	"github.com/promingy/yelp-clone/backend/internal/models"
	"github.com/promingy/yelp-clone/backend/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo    *repositories.UserRepository
	UserService *UserService
	JWTConfig   *config.JWTConfig
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	userService *UserService,
	jwtConfig *config.JWTConfig,
) *AuthService {
	return &AuthService{
		UserRepo:    userRepo,
		UserService: userService,
		JWTConfig:   jwtConfig,
	}
}

type LoginInput struct {
	Email    string
	Password string
}

type AuthResponse struct {
	AccessToken  string        `json:"access_token"`
	RefreshToken string        `json:"refresh_token"`
	User         *UserResponse `json:"user"`
}

type UserResponse struct {
	ID         int64  `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ProfilePic string `json:"profile_pic"`
}

func (s *AuthService) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	user, err := s.UserRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, fmt.Errorf("No account associated with email")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(input.Password+passwordSalt),
	); err != nil {
		return nil, fmt.Errorf("Wrong password")
	}

	profile, err := s.UserRepo.GetProfileByUserId(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve userProfile")
	}

	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate access token")
	}

	refreshToken, err := s.generateRefreshToken(user)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate refresh token")
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: &UserResponse{
			ID:         user.ID,
			Email:      user.Email,
			FirstName:  profile.FirstName,
			LastName:   profile.LastName,
			ProfilePic: profile.ProfilePic,
		},
	}, nil
}

func (s *AuthService) Register(ctx context.Context, input CreateUserInput) (*AuthResponse, error) {
	result, err := s.UserService.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.generateAccessToken(result.User)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate Access Token")
	}

	refreshToken, err := s.generateRefreshToken(result.User)
	if err != nil {
		return nil, fmt.Errorf("Frailed to generate Refresh Token")
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: &UserResponse{
			ID:         result.User.ID,
			Email:      result.User.Email,
			FirstName:  result.Profile.FirstName,
			LastName:   result.Profile.LastName,
			ProfilePic: result.Profile.ProfilePic,
		},
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshTokenString string) (*AuthResponse, error) {
	token, err := jwt.ParseWithClaims(
		refreshTokenString,
		&models.RefreshTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte(s.JWTConfig.RefreshSecret), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Invalid refresh token")
	}

	claims, ok := token.Claims.(*models.RefreshTokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid refresh token")
	}

	user, err := s.UserRepo.FindById(ctx, claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("User not found")
	}

	profile, err := s.UserRepo.GetProfileByUserId(ctx, user.ID)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve user profile")
	}

	accessToken, err := s.generateAccessToken(user)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate access token")
	}

	return &AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshTokenString,
		User: &UserResponse{
			ID:         user.ID,
			Email:      user.Email,
			FirstName:  profile.FirstName,
			LastName:   profile.LastName,
			ProfilePic: profile.ProfilePic,
		},
	}, nil
}

func (s *AuthService) ValidateAccessToken(tokenString string) (*models.AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&models.AccessTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}
			return []byte(s.JWTConfig.AccessSecret), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Invalid token: %w", err)
	}

	claims, ok := token.Claims.(*models.AccessTokenClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("Invalid token claims")
	}

	return claims, nil
}

func (s *AuthService) generateAccessToken(user *models.User) (string, error) {
	claims := &models.AccessTokenClaims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.JWTConfig.AccessExpiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.JWTConfig.AccessSecret))
}

func (s *AuthService) generateRefreshToken(user *models.User) (string, error) {
	claims := &models.RefreshTokenClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.JWTConfig.RefreshExpiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.JWTConfig.RefreshSecret))
}
