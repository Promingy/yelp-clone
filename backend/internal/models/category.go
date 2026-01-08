package models

import "github.com/uptrace/bun"

type Category struct {
	bun.BaseModel `bun:"table:categories"`

	ID       int64  `bun:"id,pk,autoincrement"`
	Category string `bun:"category"`
}
