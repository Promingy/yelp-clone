package models

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	// Exported fields become database columns
	ID           int64  `bun:"id,pk,autoincrement"`
	FirstName    string `bun:"first_name,notnull"`
	LastName     string `bun:"last_name,notnull"`
	Email        string `bun:"email,unique,notnull"`
	IsVerified   bool   `bun:"is_verified"`
	PasswordHash string `bun:"password_hash"`
	PhoneNumber  string `bun:"phone_number"`
	Bio          string `bun:"bio"`
	IsAdmin      bool   `bun:"is_admin"`
	Country      string `bun:"country"`
	State        string `bun:"state"`
	ZipCode      string `bun:"zip_code"`
	ProfilePic   string `bun:"profile_pic"`

	LastLoginAt *time.Time `bun:"last_login_at"` // Pointer so it can be null
	CreatedAt   time.Time  `bun:"created_at,notnull,default:current_timestamp"`
	UpdatedAt   time.Time  `bun:"updated_at,notnull,default:current_timestamp"`

	// Unexported fields are ignorned by bun
	password string
	cache    map[string]interface{}
}

func (u *User) Validate() map[string]string {
    errs := make(map[string]string)

    if u.FirstName == "" {
        errs["first_name"] = "First Name is required"
    }
    if u.LastName == "" {
        errs["last_name"] = "Last Name is required"
    }
    if u.Email == "" {
        errs["email"] = "Email is required"
    }

    return errs
}
