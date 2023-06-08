package orderstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order/ordermodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindOrderWithCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*ordermodel.Order, error) {
	var data ordermodel.Order
	db := s.db

	db = db.Table(ordermodel.Order{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
