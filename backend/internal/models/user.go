package models

import (
	"context"
	"strings"
	"time"

	"github.com/uptrace/bun"
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
	Password string `bun:"-" json:"-"`
	cache    map[string]interface{}
}

var _ bun.BeforeAppendModelHook = (*User)(nil)

func (u *User) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	if u.Email != "" {
		u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	}
	return nil
}
