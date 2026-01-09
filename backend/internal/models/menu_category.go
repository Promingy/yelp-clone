package models

import (
	"time"

	"github.com/uptrace/bun"
)

type MenuCategory struct {
	bun.BaseModel `bun:"table:menu_categories"`

	ID          int64     `bun:"id,pk,autoincrement"`
	MenuID      int64     `bun:"menu_id,notnull"`
	Name        string    `bun:"name,notnull"`
	Description string    `bun:"description"`
	CreatedAt   time.Time `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt   time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
}
