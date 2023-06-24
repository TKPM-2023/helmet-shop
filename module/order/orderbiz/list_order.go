package orderbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order/ordermodel"
	"context"
)

type ListOrderRepo interface {
	ListOrder(
		context context.Context,
		filter *ordermodel.Filter,
		paging *common.Paging,
	) ([]ordermodel.Order, error)
}

type listOrderBusiness struct {
	repo ListOrderRepo
}

func NewListOrderBusiness(repo ListOrderRepo) *listOrderBusiness {
	return &listOrderBusiness{repo: repo}
}

func (business *listOrderBusiness) ListOrder(context context.Context,
	filter *ordermodel.Filter,
	paging *common.Paging,
) ([]ordermodel.Order, error) {
	result, err := business.repo.ListOrder(context, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return result, nil
}
