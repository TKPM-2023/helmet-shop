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

	db = db.Table(categorymodel.Category{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i], "status = ?", 1)
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
