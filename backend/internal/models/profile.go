package models

import (
	"context"

	"github.com/uptrace/bun"
)

type Profile struct {
	bun.BaseModel `bun:"table:profiles"`

	ID     int64 `bun:"id,pk,autoincrement" json:"id"`
	UserID int64 `bun:"user_id,notnull,unique,fk:users(id)"`

	FirstName   string `bun:"first_name,notnull" json:"first_name" validate:"required,min=1,max=50"`
	LastName    string `bun:"last_name,notnull" json:"last_name" validate:"required,min=1,max=50"`
	PhoneNumber string `bun:"phone_number" json:"phone_number" validate:"omitempty,e164"`
	Bio         string `bun:"bio" json:"bio" validate:"omitempty,max=500"`
	Country     string `bun:"country" json:"country" validate:"required,iso3166_1_alpha2"`
	City        string `bun:"city" json:"city" validate:"required,min=1,max=100"`
	State       string `bun:"state" json:"state" validate:"required,min=2,max=50"`
	ZipCode     string `bun:"zip_code" json:"zip_code" validate:"required,numeric,len=5"`
	ProfilePic  string `bun:"profile_pic" json:"profile_pic" validate:"omitempty,url"`
}

var _ bun.BeforeAppendModelHook = (*Profile)(nil)

func(p *Profile) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	return nil
}