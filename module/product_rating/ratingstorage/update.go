package ratingstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingmodel"
)

func (s *sqlStore) UpdateRating(context context.Context, id int, data *ratingmodel.RatingUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
