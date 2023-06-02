package orderstorage

import (
	"TKPM-Go/module/order/ordermodel"
	"context"
)

func (s *sqlStore) UpdateOrder(context context.Context, id int, data *ordermodel.OrderUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
