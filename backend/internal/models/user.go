package models

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	// Exported fields become database columns
	ID int64 `bun:"id,pk,autoincrement" json:"id"`

	FirstName   string `bun:"first_name,notnull" json:"first_name" validate:"required,min=1,max=50"`
	LastName    string `bun:"last_name,notnull" json:"last_name" validate:"required,min=1,max=50"`
	Email       string `bun:"email,unique,notnull" json:"email" validate:"required,email"`
	PhoneNumber string `bun:"phone_number" json:"phone_number" validate:"omitempty,e164"`
	Bio         string `bun:"bio" json:"bio" validate:"omitempty,max=500"`
	Country     string `bun:"country" json:"country" validate:"required,iso3166_1_alpha2"`
	City        string `bun:"city" json:"city" validate:"required,min=1,max=100"`
	State       string `bun:"state" json:"state" validate:"required,min=2,max=50"`
	ZipCode     string `bun:"zip_code" json:"zip_code" validate:"required,numeric,len=5"`
	ProfilePic  string `bun:"profile_pic" json:"profile_pic" validate:"omitempty,url"`

	PasswordHash string `bun:"password_hash"`
	IsAdmin      bool   `bun:"is_admin"`
	IsVerified   bool   `bun:"is_verified" json:"is_verified" validate:"-"`

	LastLoginAt *time.Time `bun:"last_login_at"` // Pointer so it can be null
	CreatedAt   time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time  `bun:"updated_at,notnull,default:current_timestamp"`

	// Unexported fields are ignorned by bun
	Password string
	cache    map[string]interface{}
}

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	if u.Email != "" {
		u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	}

	if u.Password != "" && !strings.HasPrefix(u.Password, "$2a$") {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		u.PasswordHash = string(hashedPassword)
	}

	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)

	return nil
}
