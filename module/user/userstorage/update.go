package userstorage

import (
	"TKPM-Go/module/user/usermodel"
	"context"
)

func (s *sqlStore) UpdateUserInfo(context context.Context, id int, data *usermodel.UserUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
