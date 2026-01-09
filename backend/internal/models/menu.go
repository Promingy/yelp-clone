package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Menu struct {
	bun.BaseModel `bun:"table:menus"`

	ID          int64     `bun:"id,pk,autoincrement"`
	BusinessID  int64     `bun:"business_id,notnull"`
	Name        string    `bun:"name"`
	Description string    `bun:"description"`
	URL         string    `bun:"url"`
	StartDate   time.Time `bun:"start_date"`
	EndDate     time.Time `bun:"end_date"`
	CreatedAt   time.Time `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
}
