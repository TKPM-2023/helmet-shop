package productstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product/productmodel"
	"context"
)

func (s *sqlStore) FindProductsByName(ctx context.Context,
	name string,
	moreKeys ...string,
) ([]productmodel.Product, error) {
	var result []productmodel.Product
	db := s.db.Table(productmodel.Product{}.TableName())

	for i := range moreKeys {
		db.Preload(moreKeys[i], "status = ?", 1)
	}

	if err := db.Where("name LIKE ?", "%"+name+"%").Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
