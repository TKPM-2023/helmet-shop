package categorystorage

import (
	"TKPM-Go/module/category/categorymodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) UpdateCategory(context context.Context, id int, data *categorymodel.CategoryUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) IncreaseTotalProduct(context context.Context, id int) error {
	if err := s.db.Table(categorymodel.Category{}.TableName()).Where("id = ?", id).
		Update("total_product", gorm.Expr("total_product + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DecreaseTotalProduct(context context.Context, id int) error {
	if err := s.db.Table(categorymodel.Category{}.TableName()).Where("id = ?", id).
		Update("total_product", gorm.Expr("total_product - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}
