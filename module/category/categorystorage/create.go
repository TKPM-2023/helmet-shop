package categorystorage

import (
	"LearnGo/common"
	"LearnGo/module/category/categorymodel"
	"context"
)

func (s *sqlStore) CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error {
	db := s.db.Begin()
	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}
	return nil
}
