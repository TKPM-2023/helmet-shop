package productstorage

import (
	"TKPM-Go/module/product/productmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateProduct(context context.Context, id int, data *productmodel.ProductUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) IncreaseTotalRating(context context.Context, id int) error {
	if err := s.db.Table(productmodel.Product{}.TableName()).Where("id = ?", id).
		Update("total_rating", gorm.Expr("total_rating + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DecreaseTotalRating(context context.Context, id int) error {
	if err := s.db.Table(productmodel.Product{}.TableName()).Where("id = ?", id).
		Update("total_rating", gorm.Expr("total_rating - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
