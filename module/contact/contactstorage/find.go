package contactstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/contact/contactmodel"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStore) FindContactWithCondition(ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*contactmodel.Contact, error) {
	var data contactmodel.Contact
	db := s.db

	db = db.Table(contactmodel.Contact{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i], "status = ?", 1)
	}

	if err := db.Where(conditions).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
