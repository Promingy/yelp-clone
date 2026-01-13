package repositories

import (
	"context"

	"github.com/promingy/yelp-clone/backend/internal/models"
	"github.com/uptrace/bun"
)

type ProfileRepository struct {
	db *bun.DB
}

func NewProfileRepository(db *bun.DB) *ProfileRepository {
	return &ProfileRepository{db}
}

func (r *ProfileRepository) Create(ctx context.Context, profile *models.Profile) error {
	_, err := r.db.NewInsert().Model(profile).Exec(ctx)
	return err
}

func (r *ProfileRepository) FindByUserID(ctx context.Context, userID int64) (*models.Profile, error) {
	profile := new(models.Profile)
	err := r.db.NewSelect().Model(profile).Where("user_id = ?", userID).Scan(ctx)
	return profile, err
}