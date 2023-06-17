package contactstorage

import (
	"TKPM-Go/module/contact/contactmodel"
	"context"
)

func (s *sqlStore) UpdateContact(context context.Context, id int, data *contactmodel.ContactUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).
		Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
