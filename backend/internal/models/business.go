package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Business struct {
	bun.BaseModel `bun:"table:businesses,alias:b"`

	ID          int64     `bun:"id,pk,autoincrement"`
	Name        string    `bun:"name,notnull"`
	Description string    `bun:"description"`
	Email       string    `bun:"email,unique,notnull"`
	PhoneNumber string    `bun:"phone_number"`
	Website     string    `bun:"website"`
	Delivery    bool      `bun:"delivery"`

	Country     string    `bun:"country"`
	City        string    `bun:"city"`
	State       string    `bun:"state"`
	Street      string    `bun:"street"`
	ZipCode     string    `bun:"zipcode"`

	Latitude    float64   `bun:"latitude"`
	Longitude   float64   `bun:"longitude"`

	Status      string    `bun:"status,notnull,default:'pending'"`
	
	CreatedAt   time.Time `bun:"created_at"`
	UpdatedAt   time.Time `bun:"updated_at"`
}
