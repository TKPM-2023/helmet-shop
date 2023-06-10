package orderdetailstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindOrderDetailWithCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*orderdetailmodel.OrderDetail, error) {
	var data orderdetailmodel.OrderDetail
	db := s.db

	var length int64
	if err := db.Table(orderdetailmodel.OrderDetail{}.TableName()).Count(&length).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if length == 0 {
		return nil, nil
	}

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	if err := s.db.Where(conditions).First(&data).Error; err != nil {
		// case: error from DB
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
