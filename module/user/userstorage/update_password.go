package userstorage

import (
	"TKPM-Go/module/user/usermodel"
	"context"
)

func (s *sqlStore) UpdatePassword(context context.Context, id int, data *usermodel.PasswordUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
