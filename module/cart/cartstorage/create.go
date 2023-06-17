package cartstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/cart/cartmodel"
	"context"
)

func (s *sqlStore) CreateCart(ctx context.Context, data *cartmodel.CartCreate) (int, error) {
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return 0, common.ErrDB(err)
	}
	return data.Id, nil
}
