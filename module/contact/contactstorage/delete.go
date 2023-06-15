package contactstorage

import (
	"TKPM-Go/module/contact/contactmodel"
	"context"
)

func (s *sqlStore) DeleteContact(context context.Context, id int) error {
	if err := s.db.Table(contactmodel.Contact{}.TableName()).Where("id = ?", id).
		Updates(map[string]interface{}{
			"status": 0,
		}).Error; err != nil {
		return err
	}
	return nil
}
