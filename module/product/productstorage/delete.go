package productstorage

import (
	"TKPM-Go/module/product/productmodel"
	"context"
)

func (s *sqlStore) DeleteProduct(context context.Context, id int) error {
	if err := s.db.Table(productmodel.Product{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return err
	}
	return nil
}
