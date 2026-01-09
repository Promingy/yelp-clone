package models

import "github.com/uptrace/bun"

type CollectionBusiness struct {
	bun.BaseModel `bun:"table:collection_businesses"`

	CollectionId int64 `bun:"collection_id,pk,notnull"`
	BusinessID   int64 `bun:"business_id,pk,notnull"`
}
