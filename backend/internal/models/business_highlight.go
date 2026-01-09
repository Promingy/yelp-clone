package models

import "github.com/uptrace/bun"

type BusinessHighlight struct {
	bun.BaseModel `bun:"table:business_highlights"`

	BusinessID  int64 `bun:"business_id,pk,notnull"`
	HighlightID int64 `bun:"highlight_id,pk,notnull"`
}
