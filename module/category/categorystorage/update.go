package categorystorage

import (
	"TKPM-Go/module/category/categorymodel"
	"context"
)

func (s *sqlStore) UpdateCategory(context context.Context, id int, data *categorymodel.CategoryUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
