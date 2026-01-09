package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Collection struct {
	bun.BaseModel `bun:"table:collections"`

	ID          int64     `bun:"id,pk,autoincrement"`
	UserID      int64     `bun:"user_id,notnull"`
	Name        string    `bun:"name"`
	Description string    `bun:"description"`
	CreatedAt   time.Time `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
}
