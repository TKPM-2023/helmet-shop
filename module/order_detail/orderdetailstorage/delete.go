package orderdetailstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailmodel"
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
