package orderdetailstorage

import (
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"context"
)

func (s *sqlStore) UpdateOrderDetail(context context.Context, id int, data *orderdetailmodel.OrderDetailUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}