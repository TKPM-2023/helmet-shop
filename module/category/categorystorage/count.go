package categorystorage

import (
	"TKPM-Go/module/category/categorymodel"
)

func (s *sqlStore) CountCategory() int64 {
	db := s.db.Table(categorymodel.EntityName)

	var count int64
	db.Where("status=?",1).Count(&count)
	return count
}