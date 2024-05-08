package uploadstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*common.Image, error) {
	var data common.Image

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
