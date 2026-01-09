package models

import (
	"time"

	"github.com/uptrace/bun"
)

type MenuEntry struct {
	bun.BaseModel `bun:"table:menu_entries"`

	MenuId         int64     `bun:"menu_id,pk,notnull"`
	MenuItemId     int64     `bun:"menu_item_id,pk,notnull"`
	MenuCategoryID *int64    `bun:"menu_category_id"`
	Price          float64   `bun:"price,notnull"`
	IsAvailable    time.Time `bun:"is_available,notnull,default:true"`
	Special        time.Time `bun:"special,notnull,default:false"`
}
