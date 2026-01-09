package models

import "github.com/uptrace/bun"

type Highlight struct {
	bun.BaseModel `bun:"table:highlights"`

	ID        int64  `bun:"id,pk,autoincrement"`
	Highlight string `bun:"highlight"`
}
