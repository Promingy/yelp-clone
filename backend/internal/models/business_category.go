package models

import "github.com/uptrace/bun"

type BusinessCategory struct {
	bun.BaseModel `bun:"table:business_categories"`

	BusinessID int64 `bun:"business_id,pk"`
	CategoryID int64 `bun:"category_id,pk"`
}