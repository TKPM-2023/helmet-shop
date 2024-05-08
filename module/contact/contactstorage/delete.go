package contactstorage

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactmodel"
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
