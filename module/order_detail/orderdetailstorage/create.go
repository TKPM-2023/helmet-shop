package orderdetailstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"context"
)

func (s *sqlStore) CreateOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetailCreate) error {
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
