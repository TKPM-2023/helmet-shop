package orderstorage

import (
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
)

func (s *sqlStore) CountOrder() int64 {
	db := s.db.Table(ordermodel.EntityName)

	var count int64
	db.Where("status=?", 1).Count(&count)
	return count
}
