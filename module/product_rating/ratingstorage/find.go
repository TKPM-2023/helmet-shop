package ratingstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product_rating/ratingmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*ratingmodel.Rating, error) {
	var data ratingmodel.Rating
	db := s.db.Table(ratingmodel.Rating{}.TableName())

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
