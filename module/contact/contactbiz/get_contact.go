package contactbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactmodel"
)

type GetContactStore interface {
	FindContactWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*contactmodel.Contact, error)
}

type getContactBusiness struct {
	store GetContactStore
}

func NewGetContactBusiness(store GetContactStore) *getContactBusiness {
	return &getContactBusiness{store: store}
}

func (business *getContactBusiness) GetContact(
	ctx context.Context, id int) (*contactmodel.Contact, error) {
	result, err := business.store.FindContactWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(contactmodel.EntityName, err)

		}
		return nil, common.ErrCannotGetEntity(contactmodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(contactmodel.EntityName, err)
	}

	return result, err
}
