package orderrepository

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
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
