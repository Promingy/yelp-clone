package models

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	// Exported fields become database columns
	ID int64 `bun:"id,pk,autoincrement" json:"id"`

	Email       string `bun:"email,unique,notnull" json:"email" validate:"required,email"`

	PasswordHash string `bun:"password_hash,notnull" json:"-"`
	IsAdmin      bool   `bun:"is_admin"`
	IsVerified   bool   `bun:"is_verified" json:"is_verified" validate:"-"`

	LastLoginAt *time.Time `bun:"last_login_at"` // Pointer so it can be null
	CreatedAt   time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time  `bun:"updated_at,notnull,default:current_timestamp"`

	// Unexported fields are ignorned by bun
	Password string `bun:"-" validate:"required" json:"-"`
	cache    map[string]interface{}
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("No .env file found, relying on system env.")
	}
	SALT := os.Getenv("PASSWORD_SALT")

	if u.Email != "" {
		u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	}

	if u.Password != "" && !strings.HasPrefix(u.Password, "$2a$") {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password+SALT), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		u.PasswordHash = string(hashedPassword)
		u.Password = ""
	}

	return nil
}
