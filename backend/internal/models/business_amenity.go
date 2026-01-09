package models

import "github.com/uptrace/bun"

type BusinessAmenity struct {
	bun.BaseModel `bun:"table:business_amenities"`

	BusinessID int64 `bun:"business_id,pk,notnull"`
	AmenityID  int64 `bun:"amenity_id,pk,notnull"`
}
