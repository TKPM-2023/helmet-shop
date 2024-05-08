package orderstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
)

func (s *sqlStore) UpdateOrder(context context.Context, id int, data *ordermodel.OrderUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
