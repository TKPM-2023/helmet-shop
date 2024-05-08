package cartstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartmodel"
	"gorm.io/gorm"
)

func (s *sqlStore) FindCartWithCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*cartmodel.Cart, error) {
	var data cartmodel.Cart
	db := s.db.Table(cartmodel.Cart{}.TableName())

	var length int64
	if err := db.Count(&length).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if length == 0 {
		return nil, nil
	}

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
