package categorystorage

import (
	"TKPM-Go/module/category/categorymodel"
	"context"
)

func (s *sqlStore) DeleteCategory(context context.Context, id int) error {
	if err := s.db.Table(categorymodel.Category{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return err
	}
	return nil
}
