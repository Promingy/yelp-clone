package models

import (
	"time"

	"github.com/uptrace/bun"
)

type ReviewImage struct {
	bun.BaseModel `bun:"review_images"`

	ID        int64     `bun:"id,pk,autoincrement"`
	ReviewID  int64     `bun:"review_id,notnull"`
	URL       string    `bun:"url"`
	Preview   bool      `bun:"preview,default:false"`
	CreatedAt time.Time `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
}
