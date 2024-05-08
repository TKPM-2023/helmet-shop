package productstorage

import (
	"context"
	"fmt"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/product/productmodel"
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
		fmt.Println(f.Name)
		if f.Status >= 0 {
			db = db.Where("status = ?", f.Status)
		}
		if f.Name != "" {
			db = db.Where("name LIKE ?", "%"+f.Name+"%")
		}
		if f.Description != "" {
			db = db.Where("description LIKE ?", "%"+f.Description+"%")
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
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
