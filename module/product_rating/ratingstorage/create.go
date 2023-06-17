package ratingstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product_rating/ratingmodel"
	"context"
)

func (s *sqlStore) CreateRating(ctx context.Context, data *ratingmodel.RatingCreate) error {
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
