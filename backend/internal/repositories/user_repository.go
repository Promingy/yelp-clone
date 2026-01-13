package repositories

import (
	"context"

	"github.com/promingy/yelp-clone/backend/internal/models"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	user := new(models.User)
	err := r.db.NewSelect().Model(user).Where("email = ?", email).Scan(ctx)
	return user, err
}

func (r *UserRepository) CreateUserWithProfile(ctx context.Context, user *models.User, profile *models.Profile) error {
	return r.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if _, err := tx.NewInsert().Model(user).Exec(ctx); err != nil {
			return err
		}
		profile.UserID = user.ID
		if _, err := tx.NewInsert().Model(profile).Exec(ctx); err != nil {
			return err
		}
		return nil
	})
}
