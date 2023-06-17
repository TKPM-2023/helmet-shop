package orderdetailstorage

import (
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"context"
)

func (s *sqlStore) DeleteOrderDetail(context context.Context, id int) error {
	if err := s.db.Table(orderdetailmodel.OrderDetail{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return err
	}
	return nil
}
