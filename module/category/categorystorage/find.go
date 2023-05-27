package categorystorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/category/categorymodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindCategoryWithCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*categorymodel.Category, error) {
	var data categorymodel.Category
	db := s.db

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	if err := s.db.Where(conditions).First(&data).Error; err != nil {
		// case: error from DB
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
