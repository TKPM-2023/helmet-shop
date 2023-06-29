package contactbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/contact/contactmodel"
	"context"
)

type ListContactStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *contactmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]contactmodel.Contact, error)
}

type listContactBusiness struct {
	store ListContactStore
}

func NewListContactBusiness(store ListContactStore) *listContactBusiness {
	return &listContactBusiness{store: store}
}

func (business *listContactBusiness) ListContact(context context.Context,
	filter *contactmodel.Filter,
	paging *common.Paging,
) ([]contactmodel.Contact, error) {
	result, err := business.store.ListDataWithCondition(context, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(contactmodel.EntityName, err)
	}

	return result, nil
}
