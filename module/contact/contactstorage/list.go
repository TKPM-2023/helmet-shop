package contactstorage

import (
	"TKPM-Go/common"
	"TKPM-Go/module/contact/contactmodel"
	"context"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *contactmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]contactmodel.Contact, error) {
	var result []contactmodel.Contact
	db := s.db

	db = db.Table(contactmodel.Contact{}.TableName())

	if f := filter; f != nil {
		if f.Status > 0 {
			db = db.Where("status = ?", f.Status)
		}
		if f.UserId != nil {
			db = db.Where("user_id = ?", f.UserId.GetLocalID())
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for i := range moreKeys {
		db = db.Preload(moreKeys[i], "status = ?", 1)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return nil, common.ErrDB(err)
		}
		db = db.Where("id < ?", uid.GetLocalID())

	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
