package contactbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/contact/contactmodel"
	"context"
)

type CreateContactStore interface {
	CreateContact(ctx context.Context, data *contactmodel.ContactCreate) error
	FindContactWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*contactmodel.Contact, error)
}

type createContactBusiness struct {
	store CreateContactStore
}

func NewCreateContactBusiness(store CreateContactStore) *createContactBusiness {
	return &createContactBusiness{store: store}
}

func (business *createContactBusiness) CreateContact(context context.Context, data *contactmodel.ContactCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	/*
		_, err := business.store.FindContactWithCondition(context, map[string]interface{}{"user_id": data.UserId})

		if err != nil {
			return err
		}

		if result != nil {
			return common.ErrEntityExisted(contactmodel.EntityName, nil)
		}
	*/
	if err := business.store.CreateContact(context, data); err != nil {
		return common.ErrCannotCreateEntity(contactmodel.EntityName, err)
	}
	return nil
}
