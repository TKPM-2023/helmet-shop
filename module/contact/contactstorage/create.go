package contactstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/contact/contactmodel"
	"context"
)

func (s *sqlStore) CreateContact(ctx context.Context, data *contactmodel.ContactCreate) error {
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
