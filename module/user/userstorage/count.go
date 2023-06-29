package userstorage

import "TKPM-Go/module/user/usermodel"

func (s *sqlStore) CountUser() int64 {
	db := s.db.Table(usermodel.User{}.TableName())

	var count int64
	db.Where("status=?",1).Count(&count)
	return count
}
