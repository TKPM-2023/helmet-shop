package productstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/module/product/productmodel"
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
