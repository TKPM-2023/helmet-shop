package orderstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
)

func (s *sqlStore) DeleteOrder(context context.Context, id int) error {
	if err := s.db.Table(ordermodel.Order{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return err
	}
	return nil
}
