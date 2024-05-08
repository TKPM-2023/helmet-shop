package contactbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactmodel"
)

type DeleteContactStore interface {
	FindContactWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*contactmodel.Contact, error)
	DeleteContact(context context.Context, id int) error
}

type deleteContactBusiness struct {
	store DeleteContactStore
}

func NewDeleteContactBusiness(store DeleteContactStore) *deleteContactBusiness {
	return &deleteContactBusiness{store: store}
}

func (business *deleteContactBusiness) DeleteContact(context context.Context, id int) error {
	oldData, err := business.store.FindContactWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return common.ErrCannotDeleteEntity(contactmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(contactmodel.EntityName, err)
	}

	if err := business.store.DeleteContact(context, id); err != nil {
		return common.ErrCannotDeleteEntity(contactmodel.EntityName, err)
	}
	return nil
}
