package migrations

import (
	"context"

	"github.com/promingy/yelp-clone/backend/internal/models"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().
			Model((*models.BusinessCategory)(nil)).
			IfNotExists().
			Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().
			Model((*models.BusinessCategory)(nil)).
			IfExists().
			Exec(ctx)
		return err
	})
}
