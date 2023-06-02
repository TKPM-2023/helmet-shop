package orderstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order/ordermodel"
	"context"
)

func (s *sqlStore) CreateOrder(ctx context.Context, data *ordermodel.OrderCreate) error {
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
