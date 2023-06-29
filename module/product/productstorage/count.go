package productstorage

import (
	"TKPM-Go/module/product/productmodel"
)

func (s *sqlStore) CountProduct() int64 {
	db := s.db.Table(productmodel.EntityName)

	var count int64
	db.Where("status=?",1).Count(&count)
	return count
}