package models

import (
	"time"

	"github.com/uptrace/bun"
)

type MenuItemImage struct {
	bun.BaseModel `bun:"table:menu_item_image"`

	ID         int64     `bun:"id,pk,autoincrement"`
	MenuItemID int64     `bun:"menu_item_id"`
	URL        string    `bun:"url"`
	Preview    bool      `bun:"preview,default:false"`
	CreatedAt  time.Time `bun:"created_at,nullzero,default:current_timestamp"`
	UpdatedAt  time.Time `bun:"updated_at,nullzero,default:current_timestamp"`
}
