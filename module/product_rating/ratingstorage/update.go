package ratingstorage

import (
	"TKPM-Go/module/product_rating/ratingmodel"
	"context"
)

func (s *sqlStore) UpdateRating(context context.Context, id int, data *ratingmodel.RatingUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
