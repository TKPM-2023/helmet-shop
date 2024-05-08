package productstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/product/productmodel"
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
