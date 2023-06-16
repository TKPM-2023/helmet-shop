package orderbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order/ordermodel"
	"context"
)

type ListOrderStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *ordermodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]ordermodel.Order, error)
}

type listOrderBusiness struct {
	store ListOrderStore
}

func NewListOrderBusiness(store ListOrderStore) *listOrderBusiness {
	return &listOrderBusiness{store: store}
}

func (business *listOrderBusiness) ListOrder(context context.Context,
	filter *ordermodel.Filter,
	paging *common.Paging,
) ([]ordermodel.Order, error) {
	result, err := business.store.ListDataWithCondition(context, filter, paging, "Products")
	if err != nil {
		return nil, err
	}

	return result, nil
}
