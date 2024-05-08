package contactbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactmodel"
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
