package models

import (
	"time"

	"github.com/uptrace/bun"
)

type BusinessFollow struct {
	bun.BaseModel `bun:"table:businesses_followed"`

	UserID     int64 `bun:"user_id,pk,notnull"`
	BusinessID int64 `bun:"business_id,pk,notnull"`
	CreatedAt  time.Time `bun:"created_at,nullzero,default:current_timestamp"`
}
