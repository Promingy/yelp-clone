package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/promingy/yelp-clone/backend/internal/models"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		fmt.Print(" [up migration] ")
		_, err := db.NewCreateTable().Model((*models.BusinessCategory)(nil)).Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().Model((*models.BusinessCategory)(nil)).Exec(ctx)
		return err
	})
}