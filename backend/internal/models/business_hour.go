package models

import (
	"time"

	"github.com/uptrace/bun"
)

type BusinessHour struct {
	bun.BaseModel `bun:"table:business_hours"`

	ID         int64     `bun:"id,pk,autoincrement"`
	BusinessID int64     `bun:"business_id,notnull"`
	OpenTime   time.Time `bun:"open_time,notnull"`
	CloseTime  time.Time `bun:"close_time,notnull"`
	DayOfWeek  int8      `bun:"day_of_week,notnull"`
	IsClosed   bool      `bun:"is_closed,default:false"`
}
