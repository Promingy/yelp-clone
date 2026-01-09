package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/promingy/yelp-clone/backend/internal/models"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().
		Model((*models.Highlight)(nil)).
		IfNotExists().
		Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().
		Model((*models.Highlight)(nil)).
		IfExists().
		Exec(ctx)
		return err
	})
}