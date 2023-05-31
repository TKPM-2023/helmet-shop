package productstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product/productmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindProductWithCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*productmodel.Product, error) {
	var data productmodel.Product
	db := s.db

	var length int64
	if err := db.Table(productmodel.Product{}.TableName()).Count(&length).Error; err != nil {
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
