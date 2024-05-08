package orderdetailstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailmodel"
)

func (s *sqlStore) UpdateOrderDetail(context context.Context, id int, data *orderdetailmodel.OrderDetailUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
