package productstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product/productmodel"
	"context"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]productmodel.Product, error) {
	var result []productmodel.Product
	db := s.db

	db = db.Table(productmodel.EntityName)

	if f := filter; f != nil {
		if f.Status >= 0 {
			db = db.Where("status = ?", f.Status)
		}
		if f.CategoryId > 0 {
			db = db.Where("category_id = ?", f.CategoryId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i], "status = ?", 1)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())

	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(int(offset))
	}

	if err := db.
		Limit(int(paging.Limit)).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
