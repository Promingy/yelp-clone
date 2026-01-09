package models

import (
	"time"

	"github.com/uptrace/bun"
)

type BusinessImage struct {
	bun.BaseModel `bun:"table:business_images"`

	ID         int64     `bun:"id,pk,autoincrement"`
	BusinessID int64     `bun:"business_id"`
	URL        string    `bun:"url"`
	Preview    bool      `bun:"preview,default:false"`
	CreatedAt  time.Time `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt  time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
}
