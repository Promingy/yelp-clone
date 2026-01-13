package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	e "github.com/promingy/yelp-clone/backend/internal/errors"
	"github.com/promingy/yelp-clone/backend/internal/models"
	"github.com/promingy/yelp-clone/backend/internal/repositories"
	"github.com/promingy/yelp-clone/backend/internal/validation"
	v "github.com/promingy/yelp-clone/backend/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

var passwordSalt string

func init() {
	godotenv.Load()
	passwordSalt = os.Getenv("PASSWORD_SALT")
	if passwordSalt == "" {
		log.Fatal("PASSWORD_SALT not set")
	}
}

type UserService struct {
	userRepo    *repositories.UserRepository
	profileRepo *repositories.ProfileRepository
	validator   *validator.Validate
}

func NewUserService(
	userRepo *repositories.UserRepository,
	profileRepo *repositories.ProfileRepository,
	validator *validator.Validate,
) *UserService {
	return &UserService{
		userRepo,
		profileRepo,
		validator,
	}
}

type CreateUserInput struct {
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneNumber string
	Bio         string
	Country     string
	City        string
	State       string
	ZipCode     string
	ProfilePic  string
}

type CreateUserResult struct {
	User    *models.User
	Profile *models.Profile
}

func (s *UserService) CreateUser(ctx context.Context, input CreateUserInput) (*CreateUserResult, error) {
	// Validate Password
	if errs := v.ValidatePassword(input.Password); len(errs) > 0 {
		return nil, &e.ValidationError{Errors: errs}
	}

	existingUser, err := s.userRepo.FindByEmail(ctx, input.Email)
	if err == nil && existingUser != nil {
		return nil, fmt.Errorf("User with email %s already exists", input.Email)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(input.Password+passwordSalt),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	user := &models.User{
		Email:        strings.ToLower(strings.TrimSpace(input.Email)),
		PasswordHash: string(hashedPassword),
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

	if errs := validation.Validate(user); len(errs) > 0 {
		return nil, &e.ValidationError{Errors: errs}
	}
	if errs := validation.Validate(profile); len(errs) > 0 {
		return nil, &e.ValidationError{Errors: errs}
	}

	if err := s.userRepo.CreateUserWithProfile(ctx, user, profile); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &CreateUserResult{
		User:    user,
		Profile: profile,
	}, nil
}
