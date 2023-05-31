package productstorage

import (
	"TKPM-Go/module/product/productmodel"
	"context"
)

func (s *sqlStore) UpdateProduct(context context.Context, id int, data *productmodel.ProductUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
