package ratingstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingmodel"
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
