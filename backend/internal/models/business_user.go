package models

import "github.com/uptrace/bun"

type BusinessUser struct {
	bun.BaseModel `bun:"table:business_users"`

	BusinessID int64 `bun:"business_id,pk,notnull"`
	UserID int64 `bun:"user_id,pk,notnull"`
}
