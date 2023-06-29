package orderrepository

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order/ordermodel"
	"context"
)

type ListOrderStorage interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *ordermodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]ordermodel.Order, error)
}

type listOrderRepo struct {
	store ListOrderStorage
}

func NewListOrderRepo(store ListOrderStorage) *listOrderRepo {
	return &listOrderRepo{store: store}
}

func (repo *listOrderRepo) ListOrder(
	context context.Context,
	filter *ordermodel.Filter,
	paging *common.Paging,
) ([]ordermodel.Order, error) {
	result, err := repo.store.ListDataWithCondition(context, filter, paging, "Products", "Contact", "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(ordermodel.EntityName, err)
	}

	return result, nil
}
