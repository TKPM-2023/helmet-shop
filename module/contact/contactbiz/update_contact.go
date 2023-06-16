package contactbiz

import (
	"TKPM-Go/module/contact/contactmodel"
	"context"
	"errors"
)

type UpdateContactStore interface {
	UpdateContact(context context.Context, id int, data *contactmodel.ContactUpdate) error
	FindContactWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*contactmodel.Contact, error)
}

type updateContactBusiness struct {
	store UpdateContactStore
}


func NewUpdateContactBusiness(store UpdateContactStore) *updateContactBusiness {
	return &updateContactBusiness{store: store}
}

func (business *updateContactBusiness) UpdateContact(context context.Context, id int,
	data *contactmodel.ContactUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	result, err := business.store.FindContactWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}

	if result.Status == 0 {
		return errors.New("data deleted")
	}

	if err := business.store.UpdateContact(context, id, data); err != nil {
		return err
	}
	return nil

}
