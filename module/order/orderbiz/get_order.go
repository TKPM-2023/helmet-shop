package orderbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order/ordermodel"
	"context"
)

type GetOrderStore interface {
	FindOrderWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ordermodel.Order, error)
}

type getOrderBusiness struct {
	store GetOrderStore
}

func NewGetOrderBusiness(store GetOrderStore) *getOrderBusiness {
	return &getOrderBusiness{store: store}
}

func (business *getOrderBusiness) GetOrder(
	ctx context.Context, id int) (*ordermodel.Order, error) {
	result, err := business.store.FindOrderWithCondition(ctx, map[string]interface{}{"id": id}, "Products")

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(ordermodel.EntityName, err)

		}
		return nil, common.ErrCannotGetEntity(ordermodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(ordermodel.EntityName, err)
	}

	return result, err
}
