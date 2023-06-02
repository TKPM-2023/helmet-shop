package orderstorage

import (
	"TKPM-Go/module/order/ordermodel"
	"context"
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
