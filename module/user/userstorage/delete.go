package userstorage

import (
	"TKPM-Go/module/user/usermodel"
	"context"
)

func (s *sqlStore) DeleteUser(context context.Context, id int) error {
	if err := s.db.Table(usermodel.User{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return err
	}
	return nil
}
