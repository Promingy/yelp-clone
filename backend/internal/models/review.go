package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Review struct {
	bun.BaseModel `bun:"table:reviews"`

	ID         int64     `bun:"id,pk,autoincrement"`
	UserID     int64     `bun:"user_id,notnull"`
	BusinessID int64     `bun:"business_id,notnull"`
	Review     string    `bun:"review,notnull"`
	Rating     int       `bun:"rating,notnull,default:5"`
	CreatedAt  time.Time `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt  time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
}
