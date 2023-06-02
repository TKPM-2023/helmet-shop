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

type listProdcutBusiness struct {
	store ListOrderStore
}

func NewListOrderBusiness(store ListOrderStore) *listProdcutBusiness {
	return &listProdcutBusiness{store: store}
}

func (business *listProdcutBusiness) ListOrder(context context.Context,
	filter *ordermodel.Filter,
	paging *common.Paging,
) ([]ordermodel.Order, error) {
	result, err := business.store.ListDataWithCondition(context, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
