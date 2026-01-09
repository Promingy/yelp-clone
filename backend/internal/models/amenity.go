package models

import "github.com/uptrace/bun"

type Amenity struct {
	bun.BaseModel `bun:"table:amenities"`

	ID      int64  `bun:"id,pk,autoincrement"`
	Amenity string `bun:"amenity"`
}
