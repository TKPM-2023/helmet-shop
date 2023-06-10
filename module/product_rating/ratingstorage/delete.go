package ratingstorage

import (
	"TKPM-Go/module/product_rating/ratingmodel"
	"context"
)

func (s *sqlStore) DeleteRating(context context.Context, id int) error {
	if err := s.db.Table(ratingmodel.Rating{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return err
	}
	return nil
}
