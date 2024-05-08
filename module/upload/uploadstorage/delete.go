package uploadstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
)

func (s *sqlStore) DeleteImage(ctx context.Context, id int) error {
	if err := s.db.Table(common.Image{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
